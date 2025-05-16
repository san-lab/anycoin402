package facilitator

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/evmbinding"
)

func SettleHandler(c *gin.Context) {

	response := types.SettleResponse{}

	var payload Envelope

	if err := c.ShouldBindJSON(&payload); err != nil {
		reason := "error: Invalid JSON: " + err.Error()
		response.ErrorReason = &reason
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Network = payload.PaymentPayload.Network

	var from, to, tokenAddress common.Address
	var value, validAfter, validBefore *big.Int

	var nonce, r, s [32]byte
	var v byte

	from = common.HexToAddress(payload.PaymentPayload.Payload.Authorization.From)
	to = common.HexToAddress(payload.PaymentPayload.Payload.Authorization.To)
	tokenAddress = common.HexToAddress(payload.PaymentRequirements.Asset)

	// Convert value
	value, ok := new(big.Int).SetString(payload.PaymentPayload.Payload.Authorization.Value, 10)
	if !ok {
		reason := "error Invalid value format"
		response.ErrorReason = &reason
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Convert validAfter / validBefore
	validAfter, ok = new(big.Int).SetString(payload.PaymentPayload.Payload.Authorization.ValidAfter, 10)
	if !ok {
		reason := "error Invalid ValidAfter format"
		response.ErrorReason = &reason
		c.JSON(http.StatusBadRequest, response)
	}

	validBefore, ok = new(big.Int).SetString(payload.PaymentPayload.Payload.Authorization.ValidBefore, 10)
	if !ok {
		reason := "error Invalid ValidBefore format"
		response.ErrorReason = &reason
		c.JSON(http.StatusBadRequest, response)
	}

	// Convert nonce (hex string to [32]byte)
	nonceBytes, err := hex.DecodeString(strings.TrimPrefix(payload.PaymentPayload.Payload.Authorization.Nonce, "0x"))
	if err != nil || len(nonceBytes) != 32 {
		reason := "error: Invalid nonce"
		response.ErrorReason = &reason
		c.JSON(http.StatusBadRequest, response)
		return
	}
	copy(nonce[:], nonceBytes)

	// Convert r, s (hex strings to []byte)
	sig, err := hex.DecodeString(strings.TrimPrefix(payload.PaymentPayload.Payload.Signature, "0x"))
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

	privkeyhex := "56c11c2fee673894e85151857339066cd244d4932f23e660ce8502c867d0927e"
	signer, _ := crypto.HexToECDSA(privkeyhex)

	h, err := evmbinding.TransferWithAuthorization(response.Network, signer, tokenAddress, from, to, value, validAfter, validBefore, nonce, r, s, v)

	if err != nil {
		reason := fmt.Sprintf("Error parsing ABI: %s", err.Error())
		response.ErrorReason = &reason
		c.JSON(http.StatusInternalServerError, response)
	}

	response.Success = true
	response.Transaction = h.Hex()
	payer := from.Hex()
	response.Payer = &payer
	c.JSON(http.StatusOK, response)

}
