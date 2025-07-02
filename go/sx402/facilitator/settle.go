package facilitator

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"log"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	kms "github.com/proveniencenft/kmsclitool/common"
	"github.com/san-lab/sx402/all712"
	"github.com/san-lab/sx402/evmbinding"
	"github.com/san-lab/sx402/schemes"
	"github.com/san-lab/sx402/state"
)

var keyfile_name = "0xfAc178B1C359D41e9162A1A6385380de96809048.json"

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
	case schemes.Scheme_Exact_EURC, schemes.Scheme_Exact_USDC, schemes.Scheme_Exact_EUROS:
		SettleExactScheme(c, &envelope)
	case schemes.Scheme_Permit_USDC:
		SettlePermitScheme(c, &envelope)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported Scheme " + envelope.PaymentPayload.Scheme})
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
	nonceBytes, err := hex.DecodeString(strings.TrimPrefix(exactPayload.Authorization.Nonce, "0x"))
	if err != nil || len(nonceBytes) != 32 {
		reason := "error: Invalid nonce"
		response.ErrorReason = &reason
		c.JSON(http.StatusBadRequest, response)
		return
	}
	copy(nonce[:], nonceBytes)

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
		c.JSON(http.StatusInternalServerError, response)
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
	permit := new(all712.Permit)
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
	spender := permit.Message.Spender.Hex()
	response.Success = true
	response.Transaction = h.Hex()
	response.Network = envelope.PaymentRequirements.Network
	response.Payer = &spender
	c.JSON(http.StatusOK, response)

}
