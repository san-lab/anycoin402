package facilitator

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"log"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	kms "github.com/proveniencenft/kmsclitool/common"
	"github.com/san-lab/sx402/all712"
	"github.com/san-lab/sx402/evmbinding"
	"github.com/san-lab/sx402/oft"
	"github.com/san-lab/sx402/schemes"
	"github.com/san-lab/sx402/state"
)

var keyfile_name = "facilitator.json"

func InitKeys(password []byte) error {
	kf, err := kms.ReadKeyfile(keyfile_name)
	if err != nil || kf == nil {

		return fmt.Errorf("Error loading keyfile %s: %w", keyfile_name, err)
	}
	err = kf.Decrypt(password)
	if err != nil {
		log.Fatal("Error decrypting keyfile", keyfile_name, err)
	}
	fpk, err = crypto.ToECDSA(kf.Plaintext)
	if err != nil {
		return err
	}
	fmt.Println(crypto.PubkeyToAddress(fpk.PublicKey))
	keyfile = kf
	return nil
}

var keyfile *kms.Keyfile
var fpk *ecdsa.PrivateKey

func SettleHandler(c *gin.Context) {
	if fpk == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Facilitator not properly initialized"})
		c.Abort()
		return
	}

	enlp, exists := c.Get("envelope")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Envelope not found"})
		return
	}
	envelope := enlp.(all712.Envelope)

	switch envelope.PaymentPayload.Scheme {
	case schemes.Scheme_Exact_EURC, schemes.Scheme_Exact_USDC, schemes.Scheme_Exact_EURS, schemes.Scheme_Exact_Draft:
		SettleExactScheme(c, &envelope)
	case schemes.Scheme_Permit_USDC:
		SettlePermitScheme(c, &envelope)
	case schemes.Scheme_Payer0_toArbitrum, schemes.Scheme_Payer0_toBase, schemes.Scheme_Payer0M_toBase:
		SettlePayerZero(c, &envelope)
	case schemes.Scheme_Payer0Plus_toBase, schemes.Scheme_Payer0Plus_toArbitrum:
		SettleCrossChainScheme(c, &envelope)
	default:
		response := types.SettleResponse{}
		response.Network = envelope.PaymentPayload.Network
		response.Payer = &envelope.PaymentRequirements.PayTo
		response.Success = false
		reason := "Unsupported Scheme: " + envelope.PaymentPayload.Scheme
		response.ErrorReason = &reason
		c.JSON(http.StatusOK, response)
		c.Abort()
		return
	}
}

func SettleExactScheme(c *gin.Context, envelope *all712.Envelope) {

	response := types.SettleResponse{}

	exactPayload := new(types.ExactEvmPayload)
	err := json.Unmarshal(envelope.PaymentPayload.Payload, exactPayload)
	if err != nil {
		response.Success = false
		reason := fmt.Sprintf("when setting: rror unmarshalling the exact payload (%s)", err)
		response.ErrorReason = &reason
		c.JSON(http.StatusBadRequest, response)
		return
	}

	clnt, exists := c.Get("client")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No Client from middleware"})
		return
	}

	client := clnt.(*ethclient.Client)

	response.Network = envelope.PaymentPayload.Network

	var from, to, tokenAddress common.Address
	var value, validAfter, validBefore *big.Int

	var nonce, r, s [32]byte
	var v byte

	from = common.HexToAddress(exactPayload.Authorization.From)
	to = common.HexToAddress(exactPayload.Authorization.To)
	tokenAddress = common.HexToAddress(envelope.PaymentRequirements.Asset)

	// Convert value
	value, ok := new(big.Int).SetString(exactPayload.Authorization.Value, 10)
	if !ok {
		reason := "error Invalid value format"
		response.ErrorReason = &reason
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Convert validAfter / validBefore
	validAfter, ok = new(big.Int).SetString(exactPayload.Authorization.ValidAfter, 10)
	if !ok {
		reason := "error Invalid ValidAfter format"
		response.ErrorReason = &reason
		c.JSON(http.StatusBadRequest, response)
	}

	validBefore, ok = new(big.Int).SetString(exactPayload.Authorization.ValidBefore, 10)
	if !ok {
		reason := "error Invalid ValidBefore format"
		response.ErrorReason = &reason
		c.JSON(http.StatusBadRequest, response)
	}

	// Convert nonce (hex string to [32]byte)
	/*
		nonceBytes, err := hex.DecodeString(strings.TrimPrefix(exactPayload.Authorization.Nonce, "0x"))
		if err != nil || len(nonceBytes) != 32 {
			reason := "error: Invalid nonce"
			response.ErrorReason = &reason
			c.JSON(http.StatusBadRequest, response)
			return
		}
		copy(nonce[:], nonceBytes)
	*/
	nonce = common.HexToHash(exactPayload.Authorization.Nonce)

	// Convert r, s (hex strings to []byte)
	sig, err := hex.DecodeString(strings.TrimPrefix(exactPayload.Signature, "0x"))
	if err != nil || len(r) != 32 {
		reason := "error: Invalid signature format"
		response.ErrorReason = &reason
		c.JSON(http.StatusBadRequest, response)
		return
	}
	copy(r[:], sig[:32])
	copy(s[:], sig[32:64])
	v = sig[64]
	if v < 27 {
		v += 27
	}

	h, err := evmbinding.TransferWithAuthorization(client, fpk, tokenAddress, from, to, value, validAfter, validBefore, nonce, r, s, v)

	if err != nil {
		log.Println("error executing settlement", err)
		reason := fmt.Sprintf("Error parsing ABI: %s", err.Error())
		response.ErrorReason = &reason
		c.JSON(http.StatusOK, response)
		//c.Abort()
		return
	}

	state.GetReceiptCollector().Submit(*h, envelope.PaymentPayload.Network)

	response.Success = true
	response.Transaction = h.Hex()
	payer := from.Hex()
	response.Payer = &payer
	c.JSON(http.StatusOK, response)

}

func SettlePermitScheme(c *gin.Context, envelope *all712.Envelope) {
	//reuse the exact one for now
	response := types.SettleResponse{}
	permit := new(all712.PermitMessage)
	err := json.Unmarshal(envelope.PaymentPayload.Payload, permit)
	if err != nil {
		response.Success = false
		reason := fmt.Sprintf("when setting: rror unmarshalling the permit (%s)", err)
		response.ErrorReason = &reason
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = evmbinding.EnactPermit(permit, fpk)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error enacting permit": err})
		return
	}

	h, err := evmbinding.TransferFrom(permit.Message.Owner, common.HexToAddress(envelope.PaymentRequirements.PayTo),
		permit.Domain.VerifyingContract, permit.Message.Value, permit.Domain.ChainID, fpk)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error in transferFrom()": err})
		return
	}
	state.GetReceiptCollector().Submit(*h, envelope.PaymentPayload.Network)
	spender := permit.Message.Spender.Hex()
	response.Success = true
	response.Transaction = h.Hex()
	response.Network = envelope.PaymentRequirements.Network
	response.Payer = &spender
	c.JSON(http.StatusOK, response)

}

func SettlePayerZero(c *gin.Context, envelope *all712.Envelope) {

	response := types.SettleResponse{}

	clnt, exists := c.Get("client")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No Client from middleware"})
		return
	}

	pd, err := FormallyVerifyPayer0Envelope(envelope)
	if err != nil {
		reason := fmt.Sprintf("error Invalid format: %v", err)
		response.ErrorReason = &reason
		c.JSON(http.StatusBadRequest, response)
		return
	}

	client := clnt.(*ethclient.Client)

	markup, insignificant_err := evmbinding.GetMarkup(envelope.PaymentPayload.Network, envelope.PaymentRequirements.Asset, keyfile.Address)
	if insignificant_err != nil {
		log.Println(insignificant_err)
	}

	response.Network = envelope.PaymentPayload.Network
	payto, err := hex.DecodeString(envelope.PaymentRequirements.PayTo[2:])
	if err != nil {
		log.Fatalf("Error decoding payTo. This cannot happen: %v", err)
		c.JSON(http.StatusInternalServerError, "internal kms error")
		return
	}
	sendParam := new(oft.SendParam)
	sendParam.AmountLD = pd.Amount
	sendParam.MinAmountLD = big.NewInt(0).Sub(pd.Amount, markup)
	sendParam.ExtraOptions = []byte{0, 3}
	sendParam.To = [32]byte(common.LeftPadBytes(payto, 32))
	sendParam.DstEid = pd.DstEid
	sendParam.ComposeMsg = []byte{}
	sendParam.OftCmd = []byte{}

	p0token, err := oft.NewOft(pd.Asset, client)
	if err != nil {
		reason := fmt.Sprintf("error binding token contract: %v", err)
		response.ErrorReason = &reason
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Create a new CallOpts (read-only call)
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}

	// Set the _payInLzToken parameter (true or false as required).
	payInLzToken := false

	// Call quoteSend on the contract.
	messagingFee, err := p0token.QuoteSend(callOpts, *sendParam, payInLzToken)
	if err != nil {
		reason := fmt.Sprintf("error quoting send price: %v", err)
		response.ErrorReason = &reason
		c.JSON(http.StatusOK, response) //is it OK?
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(fpk, pd.chainID)
	if err != nil {
		log.Fatalf("Failed to create TransactOpts: %v", err)
		c.JSON(http.StatusInternalServerError, "internal kms error")
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		reason := fmt.Sprintf("error getting price suggestion: %v", err)
		response.ErrorReason = &reason
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Optionally customize gas
	auth.GasLimit = uint64(1000000)                           // set if needed
	auth.GasPrice = gasPrice.Add(gasPrice, big.NewInt(30000)) //
	auth.Value = messagingFee.NativeFee                       //.Mul(messagingFee.NativeFee, big.NewInt(2))

	txh, err := p0token.SendWithAuthorization(auth, *sendParam, messagingFee, pd.Payer,
		pd.ValidAfter, pd.ValidBefore, pd.nonce, pd.signature, common.HexToAddress(keyfile.Address))
	if err != nil {
		reason := fmt.Sprintf("error sending: %v", err)
		response.ErrorReason = &reason
		c.JSON(http.StatusOK, response) //is it OK?
		return

	}
	fmt.Printf("transaction hash: %s", txh.Hash().Hex())
	state.GetReceiptCollector().Submit(txh.Hash(), envelope.PaymentPayload.Network)

	response.Success = true
	response.Transaction = txh.Hash().Hex()
	response.Payer = &envelope.PaymentRequirements.PayTo
	c.JSON(http.StatusOK, response)

}
