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
	case "exact":
		VerifyExactEnvelope(c, &envelope)
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

	balance, err := evmbinding.CheckTokenBalance(client, asset, payer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Unable to check the balance": err})
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
		err = fmt.Errorf("Wrong value: %s", exactPayload.Authorization.Value)
		return
	}

	required, ok := new(big.Int).SetString(envelope.PaymentRequirements.MaxAmountRequired, 10)
	if !ok {
		err = fmt.Errorf("Wrong MaxAmountRequired value: %s", envelope.PaymentRequirements.MaxAmountRequired)
		return
	}
	if amount.Cmp(required) != 0 {
		err = fmt.Errorf("Authorized amount dirrefent from required: %v, %v", amount, required)
		return
	}

	after, ok := new(big.Int).SetString(exactPayload.Authorization.ValidAfter, 10)
	if !ok {
		err = fmt.Errorf("Wrong VelidAfter parameter: %s", exactPayload.Authorization.ValidAfter)
		return
	}

	if time.Now().Unix() < after.Int64() {
		err = fmt.Errorf("Authorization not valid yet: %v/%v", after.Int64(), time.Now().Unix())
		return
	}

	before, ok := new(big.Int).SetString(exactPayload.Authorization.ValidBefore, 10)
	if !ok {
		err = fmt.Errorf("Wrong VelidBefore parameter: %s", exactPayload.Authorization.ValidBefore)
		return
	}
	if time.Now().Unix() > before.Int64() {
		err = fmt.Errorf("Authorization expired: %v/%v", before.Int64(), time.Now().Unix())
		return
	}

	asset = common.HexToAddress(envelope.PaymentRequirements.Asset)
	chainID, ok = evmbinding.ChainIDs[envelope.PaymentPayload.Network]
	if !ok {
		err = fmt.Errorf("Unsupported network: %s", envelope.PaymentPayload.Network)
		return

	}

	einfo := ExtraInfo{}
	err = json.Unmarshal(*envelope.PaymentRequirements.Extra, &einfo)
	if err != nil {
		err = fmt.Errorf("Error parsing ExtraInfo: %w", err)
		return
	}

	ok, payer, err = all712.VerifyTransferWithAuthorizationSignature(
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
