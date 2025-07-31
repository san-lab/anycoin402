package facilitator

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/all712"
	"github.com/san-lab/sx402/evmbinding"
	"github.com/san-lab/sx402/oft"
	"github.com/san-lab/sx402/schemes"
	"github.com/san-lab/sx402/signing"
)

type ExtraInfo map[string]string

var TortugaOperator = common.HexToAddress("0xe1b783Bead4D2FDA861eA16e9D8Fa670AaD18081")

func verifyHandler(c *gin.Context) {
	enlp, exists := c.Get("envelope")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Envelope not found"})
		c.Abort()
		return
	}
	envelope := enlp.(all712.Envelope)

	switch envelope.PaymentPayload.Scheme {
	case schemes.Scheme_Exact_USDC, schemes.Scheme_Exact_EURC, schemes.Scheme_Exact_EURS:
		VerifyExactEnvelope(c, &envelope)
	case schemes.Scheme_Permit_USDC:
		VerifyPermitEnvelope(c, &envelope)
	case schemes.Scheme_Payer0_toArbitrum, schemes.Scheme_Payer0_toBase, schemes.Scheme_Payer0M_toBase:
		VerifyPayer0Envelope(c, &envelope)
	case schemes.Scheme_Payer0Plus_toBase, schemes.Scheme_Payer0Plus_toArbitrum, schemes.Scheme_Payer0Plus_toAmoy, schemes.Scheme_Payer0Plus_toOP:
		VerifyCrossChainScheme(c, &envelope)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported Scheme " + envelope.PaymentPayload.Scheme})
		c.Abort()
		return
	}
}

func VerifyExactEnvelope(c *gin.Context, envelope *all712.Envelope) {
	clnt, exists := c.Get("client")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Client not found"})
		c.Abort()
		return
	}
	client := clnt.(*ethclient.Client)

	response := types.VerifyResponse{}
	response.InvalidReason = new(string)

	parsedD, err := ParseAndVerifyExact(envelope)

	if err != nil {
		*response.InvalidReason = err.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	p := parsedD.Payer.Hex()
	response.Payer = &p
	// Checks on-chain
	reason := ""
	response.IsValid, reason = Verify3009OnChainConstraints(client, parsedD)
	response.InvalidReason = &reason
	c.JSON(http.StatusOK, response)
}

func ParseAndVerifyExact(envelope *all712.Envelope) (pd ParsedData, err error) {
	pd = ParsedData{}
	var ok bool
	exactPayload := new(types.ExactEvmPayload)
	err = json.Unmarshal(envelope.PaymentPayload.Payload, exactPayload)
	if err != nil {
		err = fmt.Errorf("error unmarshalling exact payload: %w", err)
		return
	}

	pd.Amount, ok = new(big.Int).SetString(exactPayload.Authorization.Value, 10)
	if !ok {
		err = fmt.Errorf("wrong value: %s", exactPayload.Authorization.Value)
		return
	}

	required, ok := new(big.Int).SetString(envelope.PaymentRequirements.MaxAmountRequired, 10)
	if !ok {
		err = fmt.Errorf("wrong MaxAmountRequired value: %s", envelope.PaymentRequirements.MaxAmountRequired)
		return
	}
	if pd.Amount.Cmp(required) != 0 {
		err = fmt.Errorf("authorized amount dirrefent from required: %v, %v", pd.Amount, required)
		return
	}

	pd.ValidAfter, ok = new(big.Int).SetString(exactPayload.Authorization.ValidAfter, 10)
	if !ok {
		err = fmt.Errorf("wrong VelidAfter parameter: %s", exactPayload.Authorization.ValidAfter)
		return
	}

	if time.Now().Unix() < pd.ValidAfter.Int64() {
		err = fmt.Errorf("authorization not valid yet: %v/%v", pd.ValidAfter.Int64(), time.Now().Unix())
		return
	}

	pd.ValidBefore, ok = new(big.Int).SetString(exactPayload.Authorization.ValidBefore, 10)
	if !ok {
		err = fmt.Errorf("wrong VelidBefore parameter: %s", exactPayload.Authorization.ValidBefore)
		return
	}
	if time.Now().Unix() > pd.ValidBefore.Int64() {
		err = fmt.Errorf("authorization expired: %v/%v", pd.ValidBefore.Int64(), time.Now().Unix())
		return
	}

	pd.Asset = common.HexToAddress(envelope.PaymentRequirements.Asset)
	pd.chainID, ok = evmbinding.ChainIDs[envelope.PaymentPayload.Network]
	if !ok {
		err = fmt.Errorf("unsupported network: %s", envelope.PaymentPayload.Network)
		return

	}

	if !strings.EqualFold(envelope.PaymentRequirements.PayTo, exactPayload.Authorization.To) {
		err = fmt.Errorf("destination account mismatch: %s/%s", envelope.PaymentRequirements.PayTo, exactPayload.Authorization.To)
		return
	}

	einfo := ExtraInfo{}
	err = json.Unmarshal(*envelope.PaymentRequirements.Extra, &einfo)
	if err != nil {
		err = fmt.Errorf("error parsing ExtraInfo: %w", err)
		return
	}

	pd.Payer, pd.nonce, pd.signature, err = signing.VerifyTransferWithAuthorizationSignature(
		exactPayload.Signature,
		*exactPayload.Authorization,
		einfo["name"], einfo["version"],
		pd.chainID, common.HexToAddress(envelope.PaymentRequirements.Asset))

	if err != nil {
		return
	}

	return
}

var zeroPeer [32]byte

func VerifyPayer0Envelope(c *gin.Context, envelope *all712.Envelope) {
	clnt, exists := c.Get("client")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Client not found"})
		c.Abort()
		return
	}
	client := clnt.(*ethclient.Client)

	response := types.VerifyResponse{}
	response.InvalidReason = new(string)

	pd, err := FormallyVerifyPayer0Envelope(envelope)
	if err != nil {
		*response.InvalidReason = err.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	markup, insignificant_err := evmbinding.GetMarkup(envelope.PaymentPayload.Network, envelope.PaymentRequirements.Asset, keyfile.Address)
	if insignificant_err != nil {
		log.Println(insignificant_err)
	}

	if pd.Amount.Cmp(markup) == -1 {
		err = fmt.Errorf("Slippage margin error: %v/%v", pd.Amount, markup)
		*response.InvalidReason = err.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	p := pd.Payer.Hex()
	response.Payer = &p
	// Checks on-chain
	response.IsValid, *response.InvalidReason = Verify3009OnChainConstraints(client, pd)

	if !response.IsValid {
		c.JSON(http.StatusOK, response)
		c.Abort()
		return
	}

	//TODO: Check if Peer exists
	contract, err := oft.NewOft(pd.Asset, client)
	if err != nil {
		*response.InvalidReason = fmt.Sprintf("failed to instantiate the contract: %v", err)
		response.IsValid = false
		c.JSON(http.StatusOK, response)
		c.Abort()
		return
	}
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}
	peerAtDst, err := contract.Peers(callOpts, pd.DstEid)
	if err != nil {
		*response.InvalidReason = fmt.Sprintf("error checking peers: %v", err)
		response.IsValid = false
		c.JSON(http.StatusOK, response)
		c.Abort()
		return
	}
	if peerAtDst == zeroPeer {
		*response.InvalidReason = fmt.Sprintf("No peer at dest chain: %v ", pd.DstEid)
		response.IsValid = false
		c.JSON(http.StatusOK, response)
		c.Abort()
		return
	}

	response.IsValid = true
	c.JSON(http.StatusOK, response)
}

func FormallyVerifyPayer0Envelope(envelope *all712.Envelope) (pd ParsedData, err error) {
	//payer0 envelope is an extension of the "exact", so we can reuse some code
	pd, err = ParseAndVerifyExact(envelope)
	if err != nil {
		return
	}

	payer0Payload := new(all712.Payer03009Payload)
	err = json.Unmarshal(envelope.PaymentPayload.Payload, payer0Payload)
	if err != nil {
		err = fmt.Errorf("error unmarshalling Payer03009Payload: %w", err)
		return
	}

	if payer0Payload.DestEid == 0 {
		err = fmt.Errorf("Missing destination chain id")
		return
	}
	pd.DstEid = payer0Payload.DestEid

	return
}

func Verify3009OnChainConstraints(client *ethclient.Client, pd ParsedData) (ok bool, reason string) {

	// Checks on-chain
	known, err := evmbinding.CheckAuthorizationState(client, pd.Asset, pd.Payer, pd.nonce)
	if known {
		reason = "Nonce already used"
		return
	}
	if err != nil {
		reason = fmt.Sprintf("Unable to check the auth state: %v", err)
		return
	}

	balance, err := evmbinding.CheckTokenBalance(client, pd.Asset, pd.Payer)
	if err != nil {
		reason = fmt.Sprintf("Unable to check the balance: %v", err)
		return
	}

	if pd.Amount.Cmp(balance) == 1 {
		reason = fmt.Sprintf("Insufficient balance: %v", balance)
		return
	}
	ok = true
	return
}
