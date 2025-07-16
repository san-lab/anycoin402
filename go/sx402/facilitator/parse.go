package facilitator

import (
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/all712"
	"github.com/san-lab/sx402/evmbinding"
)

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

	client, err := evmbinding.GetClientByNetwork(network)
	if err != nil {
		reason := fmt.Sprintf("could not connect to rpc: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": reason})
		return
	}
	c.Set("client", client)
	c.Next()
}

type ParsedData struct {
	Amount      *big.Int
	Markup      *big.Int
	chainID     *big.Int
	DstEid      uint32
	Payer       common.Address
	Asset       common.Address
	ValidAfter  *big.Int
	ValidBefore *big.Int
	signature   []byte
	nonce       [32]byte
}
