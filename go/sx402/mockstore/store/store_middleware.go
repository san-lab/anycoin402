package store

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/all712"
)

const X_PAYMENT_HEADER = "X-Payment"

/*
	struct {
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
*/
var EURO_SSchemaExtraBytes []byte
var USDC_SchemaExtraBytes []byte

func init() {
	var err error
	EURO_SSchemaExtraBytes, err = json.Marshal(map[string]string{"name": "EURO_S", "version": "2"})
	if err != nil {
		log.Fatal(err)
	}
	USDC_SchemaExtraBytes, err = json.Marshal(map[string]string{"name": "USDC", "version": "2"})
	if err != nil {
		log.Fatal(err)
	}

}

func NewExactEURSSchema(resource string, price string) *types.PaymentRequirements {
	return &types.PaymentRequirements{Scheme: "EURS", Network: "base-sepolia",
		PayTo:             "0xCEF702Bd69926B13ab7150624daA7aFEE0300786", //Tortuga_Governor
		MaxTimeoutSeconds: 120,
		Asset:             "0x6Ac14e603A2742fB919248D66c8ecB05D8Aec1e9",
		MaxAmountRequired: price,
		Resource:          resource,
		Extra:             (*json.RawMessage)(&EURO_SSchemaExtraBytes),
	}
}

func NewExactUSDCSchema(resource string, price string) *types.PaymentRequirements {
	return &types.PaymentRequirements{Scheme: "exact", Network: "base-sepolia",
		PayTo:             "0xCEF702Bd69926B13ab7150624daA7aFEE0300786", //Tortuga_Governor
		MaxTimeoutSeconds: 120,
		Asset:             "0x036CbD53842c5426634e7929541eC2318f3dCF7e",
		MaxAmountRequired: price,
		Resource:          resource,
		Extra:             (*json.RawMessage)(&USDC_SchemaExtraBytes),
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
	resourceURI := fmt.Sprintf("%s/resource?RESID=%s", StorePrefix, rid)
	schemaEURS := NewExactEURSSchema(resourceURI, price)
	europrice, _ := strconv.Atoi(price)
	usdprice := europrice * 11 / 10
	schemaUSDC := NewExactUSDCSchema(resourceURI, fmt.Sprintf("%v", usdprice))

	if paymentHeader == "" {

		response := gin.H{
			"x402Version": 1,
			"error":       "X-PAYMENT header is required",
			"accepts":     []any{schemaEURS, schemaUSDC},
		}
		c.JSON(http.StatusPaymentRequired, response)
		c.Abort()
		return
	}

	var env = new(all712.Envelope)
	headerPayload := new(types.PaymentPayload)
	if err := json.Unmarshal([]byte(paymentHeader), &headerPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad X-Payment header"})
	}
	var schema *types.PaymentRequirements
	if headerPayload.Scheme == "EURS" {
		schema = schemaEURS
	} else {
		schema = schemaUSDC
	}

	env.X402Version = 1
	env.PaymentRequirements = schema
	env.PaymentPayload = headerPayload

	if err := validatePayment(env); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "Invalid or unverified payment",
			"details": err.Error(),
		})
		c.Abort()
		return
	}

	settleResponse, err := settlePayment(env)
	if err != nil {
		c.JSON(http.StatusPaymentRequired, gin.H{"error settling payment": err})
		c.Abort()
		return
	}

	if settleResponse.Success {
		c.Set("settleReponse", settleResponse)
		c.Next()
	}

	c.JSON(http.StatusForbidden, gin.H{
		"settling error": settleResponse.ErrorReason,
	})
	c.Abort()
	return

}

const facilitatorURI = "https://anycoin402.duckdns.org/facilitator"

// const facilitatorURI = "http://localhost:3010/facilitator"
//const facilitatorURI = "https://x402.org/facilitator"

func validatePayment(env *all712.Envelope) error {
	// Step 1: Parse the payment header

	reqBody, err := json.Marshal(env)
	if err != nil {
		return fmt.Errorf("failed to encode request: %w", err)
	}

	resp, err := http.Post(facilitatorURI+"/verify", "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return fmt.Errorf("facilitator error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("facilitator rejected payment: %s", resp.Status)
	}

	fvres := new(types.VerifyResponse)
	err = json.NewDecoder(resp.Body).Decode(fvres)
	if err != nil {
		return fmt.Errorf("error paring reponse from the facilitator/verify: %w", err)
	}

	if fvres.IsValid {
		return nil
	}

	return fmt.Errorf("Authorization validation failed: %ss", *fvres.InvalidReason)
}

func settlePayment(env *all712.Envelope) (*types.SettleResponse, error) {
	reqBody, err := json.Marshal(env)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request: %w", err)
	}

	resp, err := http.Post(facilitatorURI+"/settle", "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("facilitator error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("facilitator rejected procesing the payment: %s", resp.Status)
	}

	stres := new(types.SettleResponse)
	err = json.NewDecoder(resp.Body).Decode(stres)
	if err != nil {
		return nil, fmt.Errorf("error paring reponse from the facilitator/settle: %w", err)
	}
	return stres, nil
}
