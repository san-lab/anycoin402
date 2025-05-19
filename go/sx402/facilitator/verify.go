package facilitator

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/all712"
	"github.com/san-lab/sx402/evmbinding"
)

type ExtraInfo map[string]string

func ParseEnvelope(c *gin.Context) {
	log.Println("in middleware")
	var payload all712.Envelope
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	if payload.X402Version == 0 {
		reason := "Empty envelope/nil Version"
		c.JSON(http.StatusBadRequest, gin.H{"error": reason})
	}
	c.Set("envelope", payload)
	c.Next()
}
func SetupClient(c *gin.Context) {
	enlp, exists := c.Get("envelope")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Envelope not found"})
		return
	}
	payload := enlp.(all712.Envelope)

	network := payload.PaymentPayload.Network
	if len(network) == 0 {
		reason := "No network specified"
		c.JSON(http.StatusBadRequest, gin.H{"error": reason})
		return
	}
	url, ok := evmbinding.RpcEndpoints[network]
	if !ok {
		reason := "Unsupported network: " + network
		c.JSON(http.StatusBadRequest, gin.H{"error": reason})
		return
	}
	client, err := ethclient.Dial(url)
	if err != nil {
		reason := fmt.Sprintf("could not connect to rpc: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": reason})
		return
	}
	c.Set("client", client)
	c.Next()
}

func verifyHandler(c *gin.Context) {
	enlp, exists := c.Get("envelope")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Envelope not found"})
		return
	}
	payload := enlp.(all712.Envelope)

	clnt, exists := c.Get("client")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Envelope not found"})
		return
	}
	client := clnt.(*ethclient.Client)

	response := types.VerifyResponse{}

	sig := payload.PaymentPayload.Payload.Signature
	auth := payload.PaymentPayload.Payload.Authorization
	einfo := ExtraInfo{}
	err := json.Unmarshal(*payload.PaymentRequirements.Extra, &einfo)
	if err != nil {
		reason := "Invalid ExtraInfo"
		response.InvalidReason = &reason
		c.JSON(http.StatusBadRequest, response)
		return
	}
	asset := common.HexToAddress(payload.PaymentRequirements.Asset)
	amount, ok := new(big.Int).SetString(payload.PaymentPayload.Payload.Authorization.Value, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Authorization Value: " + err.Error()})
		return
	}
	chainID, ok := all712.ChainIDs[payload.PaymentPayload.Network]
	if !ok {
		reason := "Unsupported network: " + payload.PaymentPayload.Network
		response.InvalidReason = &reason
		c.JSON(http.StatusOK, reason)
	}
	ok, payer, err := all712.VerifyTransferWithAuthorizationSignature(sig, *auth, einfo["name"], einfo["version"],
		chainID, asset)
	ph := payer.Hex()
	response.Payer = &ph

	if !ok || err != nil {
		reason := err.Error()
		response.InvalidReason = &reason
		c.JSON(http.StatusOK, response)
		return
	}

	noncesl, _ := hex.DecodeString(payload.PaymentPayload.Payload.Authorization.Nonce[2:])
	var nonce [32]byte
	copy(nonce[:], noncesl)

	known, err := evmbinding.CheckAuthorizationState(client, asset, payer, nonce)
	if known {
		reason := "Nonce already used"
		response.InvalidReason = &reason
		c.JSON(http.StatusOK, response)
		return
	}

	balance, err := evmbinding.CheckTokenBalance(client, asset, payer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Formal verification": ok,
			"Unable to check the balance": err,
		})
		return
	}

	if amount.Cmp(balance) == 0 {
		reason := fmt.Sprintf("Insufficient balance: %v", balance)
		response.InvalidReason = &reason
		c.JSON(http.StatusOK, response)
		return
	}

	response.IsValid = true
	c.JSON(http.StatusOK, response)
}
