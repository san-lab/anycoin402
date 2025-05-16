package facilitator

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
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
	response := types.VerifyResponse{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}
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

	ok, payer, err := all712.VerifyTransferWithAuthorizationSignature(sig, *auth, einfo["name"], einfo["version"],
		big.NewInt(84532), asset)
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

	known, err := evmbinding.CheckAuthorizationState(payload.PaymentRequirements.Network, asset, payer, nonce)
	if known {
		reason := "Nonce already used"
		response.InvalidReason = &reason
		c.JSON(http.StatusOK, response)
		return
	}

	balance, err := evmbinding.CheckTokenBalance(payload.PaymentPayload.Network, asset, payer)
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

func Start() {
	router := gin.Default()

	// Define the endpoint
	router.GET("/facilitator/*action", func(c *gin.Context) {
		c.Writer.WriteString("You probably want to use POST method for your action: " + c.Param("action"))
	})
	router.POST("/facilitator/verify", verifyHandler)
	router.GET("/", func(c *gin.Context) {
		c.Writer.WriteString("Hello there!")
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Route not found", "path": c.Request.RequestURI,
		})
	})

	// Start the server
	router.Run(":3010")
}
