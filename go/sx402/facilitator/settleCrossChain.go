package facilitator

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/all712"
	"github.com/san-lab/sx402/evmbinding"
	"github.com/san-lab/sx402/oftcc"
	"github.com/san-lab/sx402/state"
)

func SettleCrossChainScheme(c *gin.Context, envelope *all712.Envelope) {

	response := types.SettleResponse{}

	clnt, exists := c.Get("client")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No Client from middleware"})
		return
	}
	client := clnt.(*ethclient.Client)

	ccmsg, _, err := parseCrossChainMessage(envelope)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Parsing error": err.Error()})
		c.Abort()
		return
	}

	markup, insignificant_err := evmbinding.GetDetailedMarkup(
		envelope.PaymentPayload.Network,
		envelope.PaymentRequirements.Asset,
		uint32(ccmsg.Authorization.DestinationChain.Uint64()),
		keyfile.Address)
	if insignificant_err != nil {
		log.Println(insignificant_err)
	}
	//The authorzed amount must cover the slippage and the minAmount
	//This is redundant, as already checked in /verify, but somehow feels needed

	if new(big.Int).Sub(ccmsg.Authorization.Amount, markup).Cmp(ccmsg.Authorization.MinimalAmount) < 0 {
		reason := "minAmount not guaranteed"
		response.Success = false
		response.ErrorReason = &reason
		c.JSON(http.StatusOK, response)
	}

	response.Network = envelope.PaymentPayload.Network
	payto := ccmsg.Authorization.To
	sendParam := new(oftcc.SendParam)
	sendParam.AmountLD = ccmsg.Authorization.Amount
	sendParam.MinAmountLD = ccmsg.Authorization.MinimalAmount
	sendParam.ExtraOptions = []byte{0, 3}
	sendParam.To = [32]byte(common.LeftPadBytes(payto.Bytes(), 32))
	sendParam.DstEid = uint32(ccmsg.Authorization.DestinationChain.Uint64())
	sendParam.ComposeMsg = []byte{}
	sendParam.OftCmd = []byte{}

	p0token, err := oftcc.NewOftcc(ccmsg.Domain.VerifyingContract, client)
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

	auth, err := bind.NewKeyedTransactorWithChainID(fpk, ccmsg.Domain.ChainID)
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

	sig := common.Hex2Bytes(strings.TrimPrefix(ccmsg.Signature, "0x"))

	txh, err := p0token.SendWithCCAuthorization(
		auth,
		*sendParam,
		messagingFee,
		ccmsg.Authorization.From,
		ccmsg.Authorization.ValidAfter,
		ccmsg.Authorization.ValidBefore,
		common.HexToHash(ccmsg.Authorization.Nonce),
		sig,
		common.HexToAddress(keyfile.Address))
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
