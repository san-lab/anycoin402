package facilitator

import (
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/all712"
	"github.com/san-lab/sx402/evmbinding"
)

type Envelope struct {
	X402Version         int                       `json:"x402Version"`
	PaymentPayload      types.PaymentPayload      `json:paymentPayload`
	PaymentRequirements types.PaymentRequirements `json:paymentRequirements`
}

type ExtraInfo map[string]string

func verifyHandler(c *gin.Context) {
	var payload Envelope

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}
	sig := payload.PaymentPayload.Payload.Signature
	auth := payload.PaymentPayload.Payload.Authorization
	einfo := ExtraInfo{}
	err := json.Unmarshal(*payload.PaymentRequirements.Extra, &einfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Extra info: " + err.Error()})
		return
	}
	asset := common.HexToAddress(payload.PaymentRequirements.Asset)
	amount, ok := new(big.Int).SetString(payload.PaymentPayload.Payload.Authorization.Value, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Authorization Value: " + err.Error()})
		return
	}

	ok, payer, err := all712.VerifyTransferWithAuthorizationSignature(sig, *auth, einfo["name"], einfo["version"],
		big.NewInt(84532), asset)

	balance, err := evmbinding.CheckTokenBalance(payload.PaymentPayload.Network, asset, payer)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Formal verification": ok,
			"Unable to check the balance": err,
		})
		return
	}

	if amount.Cmp(balance) == 0 {
		c.JSON((http.StatusNotAcceptable), gin.H{"Insufficient balance": balance})
		return
	}

	// You can now use `payload` here
	c.JSON(http.StatusOK, gin.H{
		"message":   "Payload received",
		"signature": ok,
		"payer":     payer,
		"balance":   "sufficient",
		"nonce":     "fresh",
	})
}

func Start() {
	router := gin.Default()

	// Define the endpoint
	router.POST("/facilitator/verify", verifyHandler)
	router.GET("/", func(c *gin.Context) {
		c.Writer.WriteString("Hello there!")
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Route not found",
		})
	})

	// Start the server
	router.Run("3010")
}
