package store

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/all712"
	"github.com/san-lab/sx402/evmbinding"
	"github.com/san-lab/sx402/schemes"
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

	ac := Accepts{}

	ac.addRequirement(schemes.Scheme_Exact_USDC, evmbinding.Base_sepolia, resourceURI, usdprice)
	ac.addRequirement(schemes.Scheme_Exact_USDC, evmbinding.Amoy, resourceURI, usdprice)
	/*
		addRequirement(schemes.Scheme_Exact_EURS, evmbinding.Base_sepolia, resourceURI, price, &accepts)
		addRequirement(schemes.Scheme_Exact_EURS, evmbinding.Arbitrum_sepolia, resourceURI, price, &accepts)
		addRequirement(schemes.Scheme_Payer0_toArbitrum, evmbinding.Base_sepolia, resourceURI, price, &accepts)
		addRequirement(schemes.Scheme_Payer0_toBase, evmbinding.Arbitrum_sepolia, resourceURI, price, &accepts)
		addRequirement(schemes.Scheme_Permit_USDC, evmbinding.Base_sepolia, resourceURI, usdprice, &accepts)
		//HACK
	*/

	//try to get markup

	priceWithMarkup := GetPriceWithMarkupAsString(europrice, schemes.Scheme_Payer0Plus_toBase, evmbinding.Arbitrum_sepolia, "40245")

	ac.addRequirement(schemes.Scheme_Payer0Plus_toBase, evmbinding.Arbitrum_sepolia, resourceURI, priceWithMarkup)

	prWiM := GetPriceWithMarkupAsString(europrice, schemes.Scheme_Payer0Plus_toArbitrum, evmbinding.Base_sepolia, "40231")
	//ac.addRequirement(schemes.Scheme_Payer0Plus_toArbitrum, evmbinding.Base_sepolia, resourceURI, prWiM)
	ac.addSchemeInstance(schemes.P0_Base_toArbitrum, resourceURI, prWiM)

	ac.addSchemeInstance(schemes.P0_OP_toBase, resourceURI, price)

	if paymentHeader == "" {

		response := gin.H{
			"x402Version": 1,
			"error":       "X-PAYMENT header is required",
			"accepts":     ac,
		}
		c.JSON(http.StatusPaymentRequired, response)
		c.Abort()
		return
	}

	var env = new(all712.Envelope)
	headerPayload, err := unmarshallXPaymentHeader(paymentHeader)
	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{"error": "Bad X-Payment header", "details": err})
		c.Abort()
		return
	}

	for _, req := range ac {
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
	env.PaymentPayload = &all712.PaymentPayload{
		X402Version: headerPayload.X402Version,
		Scheme:      headerPayload.Scheme,
		Network:     headerPayload.Network,
		Payload:     headerPayload.Payload,
	}

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
		c.Set("network", headerPayload.Network)
		explorer := evmbinding.ExplorerURLs[headerPayload.Network]
		if strings.HasPrefix(headerPayload.Scheme, "payer0") || strings.HasPrefix(headerPayload.Scheme, "PZ_") {
			explorer = "https://testnet.layerzeroscan.com/"
		}
		c.Set("explorer", explorer)
		c.Set("facilitator", facilitatorURI)
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

type MarkupResponse struct {
	Markup string `json:"markup"`
}

func unmarshallXPaymentHeader(header string) (ppld *all712.PaymentPayload, err error) {
	ppld = new(all712.PaymentPayload)
	var headerbts []byte
	if !strings.Contains(header, "{") { //Not pure json, let us try base64
		headerbts, err = base64.StdEncoding.DecodeString(header)
		if err != nil {
			err = fmt.Errorf("wrong X-Payment header: %s", header)
			return
		}
	} else {
		headerbts = []byte(header)
	}
	err = json.Unmarshal(headerbts, ppld)
	return
}

func GetPriceWithMarkupAsString(price int, scheme_name, network, dstEid string) string {
	markup := 0
	markupQuery := fmt.Sprintf("%s/markup?scheme=%s&network=%s", facilitatorURI, scheme_name, network)
	if len(dstEid) > 0 {
		markupQuery += "&dstEid=" + dstEid
	}
	resp, err := http.Get(markupQuery)
	if err != nil {
		log.Println(err)
	} else {
		var result MarkupResponse
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			fmt.Println(err)
		} else {
			markup, err = strconv.Atoi(result.Markup)
			if err != nil {
				log.Println(err)
			}
		}
		price += markup

	}
	return fmt.Sprintf("%v", price)
}
