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
	"github.com/san-lab/sx402/evmbinding"
	"github.com/san-lab/sx402/schemas"
)

const X_PAYMENT_HEADER = "X-Payment"

const store_wallet = "0xCEF702Bd69926B13ab7150624daA7aFEE0300786"

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

	europrice, _ := strconv.Atoi(price)
	usdpricei := europrice * 11 / 10
	usdprice := fmt.Sprintf("%v", usdpricei)

	accepts := []*types.PaymentRequirements{}

	network := evmbinding.Base_sepolia
	usdcs, err := schemas.GetSchema("exact", network)
	if err != nil {
		log.Println(err)

	} else {

		accepts = append(accepts, usdcs.Requirement(resourceURI, usdprice, store_wallet))
	}
	euros, err := schemas.GetSchema("EUROS", network)
	if err != nil {
		log.Println(err)

	} else {
		accepts = append(accepts, euros.Requirement(resourceURI, price, store_wallet))
	}

	amoyusdc, err := schemas.GetSchema("exact", evmbinding.Amoy)
	if err != nil {
		log.Println(err)

	} else {
		accepts = append(accepts, amoyusdc.Requirement(resourceURI, usdprice, store_wallet))
	}

	sepoliaeurc, err := schemas.GetSchema("EURC", evmbinding.Sepolia)
	if err != nil {
		log.Println(err)

	} else {
		accepts = append(accepts, sepoliaeurc.Requirement(resourceURI, price, store_wallet))
	}

	zksyncussdc, err := schemas.GetSchema("exact", evmbinding.ZkSync_sepolia)
	if err != nil {
		log.Println(err)

	} else {
		accepts = append(accepts, zksyncussdc.Requirement(resourceURI, usdprice, store_wallet))
	}

	if paymentHeader == "" {

		response := gin.H{
			"x402Version": 1,
			"error":       "X-PAYMENT header is required",
			"accepts":     accepts,
		}
		c.JSON(http.StatusPaymentRequired, response)
		c.Abort()
		return
	}

	var env = new(all712.Envelope)
	headerPayload := new(types.PaymentPayload)
	if err := json.Unmarshal([]byte(paymentHeader), &headerPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad X-Payment header"})
		c.Abort()
		return
	}

	for _, req := range accepts {
		if req.Network == headerPayload.Network && req.Scheme == headerPayload.Scheme {
			env.PaymentRequirements = req
			break
		}
	}
	if env.PaymentRequirements == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad X-Payment header"})
		c.Abort()
		return
	}

	env.X402Version = 1
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
		c.JSON(http.StatusPaymentRequired, gin.H{"error": "error settling the payment", "details": err.Error()})
		c.Abort()
		return
	}

	if settleResponse.Success {
		c.Set("settleReponse", settleResponse)
		explorer := evmbinding.ExplorerURLs[headerPayload.Network]
		c.Set("explorer", explorer)
		c.Next()
	} else {

		c.JSON(http.StatusForbidden, gin.H{
			"error": settleResponse.ErrorReason,
		})
		c.Abort()
		return
	}

}

//const facilitatorURI = "https://anycoin402.duckdns.org/facilitator"

const facilitatorURI = "http://localhost:3010/facilitator"

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
		return nil, fmt.Errorf("facilitator rejected processing the payment: %s", resp.Status)
	}

	stres := new(types.SettleResponse)
	err = json.NewDecoder(resp.Body).Decode(stres)
	if err != nil {
		return nil, fmt.Errorf("error paring reponse from the facilitator/settle: %w", err)
	}
	return stres, nil
}
