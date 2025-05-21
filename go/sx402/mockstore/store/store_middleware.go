package store

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/all712"
)

const X_PAYMENT_HEADER = "X-Payment"

type Schema struct {
	Scheme            string            `json:"scheme"`
	Network           string            `json:"network"`
	MaxAmountRequired string            `json:"maxAmountRequired"`
	Resource          string            `json:"resource"`
	Description       string            `json:"description"`
	MimeType          string            `json:"mimeType"`
	PayTo             string            `json:"payTo"`
	MaxTimeoutSeconds int               `json:"maxTimeoutSeconds"`
	Asset             string            `json:"asset"`
	Extra             map[string]string `json:"extra"`
}

func NewExactEURSSchema(resource string, price string) Schema {
	return Schema{Scheme: "exact", Network: "base-sepolia",
		PayTo:             "0xCEF702Bd69926B13ab7150624daA7aFEE0300786", //Tortuga_Governor
		MaxTimeoutSeconds: 120,
		Asset:             "0x6Ac14e603A2742fB919248D66c8ecB05D8Aec1e9",
		MaxAmountRequired: price,
		Resource:          resource,
		Extra:             map[string]string{"name": "EURO_S", "version": "2"},
	}
}

var prices = map[string]string{"1": "1000", "2": "2000", "3": "3000"}

func X402Middleware(c *gin.Context) {
	rid := c.Query("RESID")
	if len(rid) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No resource specified"})
		c.Abort()
		return
	}

	price, ok := prices[rid]

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid resource ID: " + rid})
		c.Abort()
		return
	}

	paymentHeader := c.GetHeader(X_PAYMENT_HEADER)

	if paymentHeader == "" {
		resourceURI := fmt.Sprintf("/resource?RESID=%s", rid)
		schema := NewExactEURSSchema(resourceURI, price)
		response := gin.H{
			"x402Version": 1,
			"error":       "X-PAYMENT header is required",
			"accepts":     []any{schema},
		}
		c.JSON(http.StatusPaymentRequired, response)
		c.Abort()
		return
	}

	if err := validatePayment(paymentHeader); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "Invalid or unverified payment",
			"details": err.Error(),
		})
		c.Abort()
		return
	}

	// Payment is valid
	c.Next()
}

const facilitatorURI = "https://anycoin402.duckdns.org/facilitator"

func validatePayment(paymentHeader string) error {
	// Step 1: Parse the payment header
	var env all712.Envelope
	if err := json.Unmarshal([]byte(paymentHeader), &env); err != nil {
		return fmt.Errorf("failed to parse Payment header: %w", err)
	}

	// Optional: Validate content fields here
	if env.PaymentPayload == nil || env.PaymentRequirements == nil {
		return fmt.Errorf("incomplete payment envelope")
	}

	reqBody, err := json.Marshal(gin.H{"payment": paymentHeader})
	if err != nil {
		return fmt.Errorf("failed to encode request: %w", err)
	}

	resp, err := http.Post(facilitatorURI, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return fmt.Errorf("facilitator error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("facilitator rejected payment: %s", resp.Status)
	}

	return nil
}
