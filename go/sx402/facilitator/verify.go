package facilitator

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/all712"
	"github.com/san-lab/sx402/evmbinding"
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
	case schemes.Scheme_Exact_USDC, schemes.Scheme_Exact_EURC, schemes.Scheme_Exact_EUROS:
		VerifyExactEnvelope(c, &envelope)
	case schemes.Scheme_Permit_USDC:
		VerifyPermitEnvelope(c, &envelope)
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

	amount, _, payer, asset, nonce, err := FormallyVerifyExactEnvelope(envelope)

	if err != nil {
		reason := err.Error()
		response.InvalidReason = &reason
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//Diagnostic blacklist
	if payer.Hex() == TortugaOperator.Hex() {
		reason := "Tortuga Operator has been blacklisted"
		response.InvalidReason = &reason
		c.JSON(http.StatusOK, response)
		return
	}

	p := payer.Hex()
	response.Payer = &p
	// Checks on-chain
	known, err := evmbinding.CheckAuthorizationState(client, asset, payer, nonce)
	if known {
		reason := "Nonce already used"
		response.InvalidReason = &reason
		c.JSON(http.StatusOK, response)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Unable to check the auth state": err})
		c.Abort()
		return
	}

	balance, err := evmbinding.CheckTokenBalance(client, asset, payer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Unable to check the balance": err})
		c.Abort()
		return
	}

	if amount.Cmp(balance) == 1 {
		reason := fmt.Sprintf("Insufficient balance: %v", balance)
		response.InvalidReason = &reason
		c.JSON(http.StatusOK, response)
		return
	}

	response.IsValid = true
	c.JSON(http.StatusOK, response)
}

func FormallyVerifyExactEnvelope(envelope *all712.Envelope) (amount, chainID *big.Int, payer, asset common.Address, nonce [32]byte, err error) {
	var ok bool
	exactPayload := new(types.ExactEvmPayload)
	err = json.Unmarshal(envelope.PaymentPayload.Payload, exactPayload)
	if err != nil {
		err = fmt.Errorf("error unmarshalling exact payload: %w", err)
		return
	}

	amount, ok = new(big.Int).SetString(exactPayload.Authorization.Value, 10)
	if !ok {
		err = fmt.Errorf("wrong value: %s", exactPayload.Authorization.Value)
		return
	}

	required, ok := new(big.Int).SetString(envelope.PaymentRequirements.MaxAmountRequired, 10)
	if !ok {
		err = fmt.Errorf("wrong MaxAmountRequired value: %s", envelope.PaymentRequirements.MaxAmountRequired)
		return
	}
	if amount.Cmp(required) != 0 {
		err = fmt.Errorf("authorized amount dirrefent from required: %v, %v", amount, required)
		return
	}

	after, ok := new(big.Int).SetString(exactPayload.Authorization.ValidAfter, 10)
	if !ok {
		err = fmt.Errorf("wrong VelidAfter parameter: %s", exactPayload.Authorization.ValidAfter)
		return
	}

	if time.Now().Unix() < after.Int64() {
		err = fmt.Errorf("authorization not valid yet: %v/%v", after.Int64(), time.Now().Unix())
		return
	}

	before, ok := new(big.Int).SetString(exactPayload.Authorization.ValidBefore, 10)
	if !ok {
		err = fmt.Errorf("wrong VelidBefore parameter: %s", exactPayload.Authorization.ValidBefore)
		return
	}
	if time.Now().Unix() > before.Int64() {
		err = fmt.Errorf("authorization expired: %v/%v", before.Int64(), time.Now().Unix())
		return
	}

	asset = common.HexToAddress(envelope.PaymentRequirements.Asset)
	chainID, ok = evmbinding.ChainIDs[envelope.PaymentPayload.Network]
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

	payer, err = signing.VerifyTransferWithAuthorizationSignature(
		exactPayload.Signature,
		*exactPayload.Authorization,
		einfo["name"], einfo["version"],
		chainID, common.HexToAddress(envelope.PaymentRequirements.Asset))

	if err != nil {
		return
	}
	var noncesl []byte
	noncesl, err = hex.DecodeString(strings.TrimPrefix(exactPayload.Authorization.Nonce, "0x"))
	copy(nonce[:], noncesl)

	return
}

func VerifyPermitEnvelope(c *gin.Context, envelope *all712.Envelope) {
	clnt, exists := c.Get("client")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Client not found"})
		c.Abort()
		return
	}
	client := clnt.(*ethclient.Client)

	response := types.VerifyResponse{}

	permit, err := FormallyVerifyPermitEnvelope(envelope)
	if err != nil {
		reason := err.Error()
		response.InvalidReason = &reason
		c.JSON(http.StatusBadRequest, response)
		return
	}

	p := permit.Message.Owner.Hex()
	response.Payer = &p
	// Checks on-chain

	//check if nonce(s) is correct

	balance, err := evmbinding.CheckTokenBalance(client, permit.Domain.VerifyingContract, permit.Message.Owner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Unable to check the balance": err})
		return
	}

	if permit.Message.Value.Cmp(balance) == 1 {
		reason := fmt.Sprintf("Insufficient balance: %v", balance)
		response.InvalidReason = &reason
		c.JSON(http.StatusOK, response)
		return
	}

	response.IsValid = true
	c.JSON(http.StatusOK, response)
}

func FormallyVerifyPermitEnvelope(envelope *all712.Envelope) (permit *all712.Permit, err error) {
	var ok bool
	permit = new(all712.Permit)
	err = json.Unmarshal(envelope.PaymentPayload.Payload, permit)
	if err != nil {
		err = fmt.Errorf("error unmarshalling permit payload: %w", err)
		return
	}

	//This may needa more systematic design
	extraInfo := new(ExtraInfo)
	err = json.Unmarshal(*envelope.PaymentRequirements.Extra, extraInfo)
	if err != nil {
		err = fmt.Errorf("error unmarshalling extra permit info")
		return
	}
	eFacilitator, ok := (*extraInfo)["facilitator"]

	if !ok || !strings.EqualFold(eFacilitator, keyfile.Address) {
		err = fmt.Errorf("missing or wrong faclitator: %s", eFacilitator)
		return
	}

	amount := permit.Message.Value
	if amount == nil {
		err = fmt.Errorf("nil value in permit message")
		return
	}

	required, ok := new(big.Int).SetString(envelope.PaymentRequirements.MaxAmountRequired, 10)
	if !ok {
		err = fmt.Errorf("wrong MaxAmountRequired value: %s", envelope.PaymentRequirements.MaxAmountRequired)
		return
	}
	if amount.Cmp(required) != 0 {
		err = fmt.Errorf("authorized amount dirrefent from required: %v, %v", amount, required)
		return
	}

	if time.Now().Unix() > permit.Message.Deadline.Int64() {
		err = fmt.Errorf("authorization expired: %v/%v", permit.Message.Deadline.Uint64(), time.Now().Unix())
		return
	}

	chainID, ok := evmbinding.ChainIDs[envelope.PaymentPayload.Network]
	if !ok {
		err = fmt.Errorf("unsupported network: %s", envelope.PaymentPayload.Network)
		return

	}

	if chainID.Cmp(permit.Domain.ChainID) != 0 {
		err = fmt.Errorf("ChainID mismatch: %v/%v", chainID, permit.Domain.ChainID)
		return
	}

	_, err = signing.VerifyPermitSignature(permit)
	return

}
