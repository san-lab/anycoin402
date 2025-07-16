package facilitator

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/all712"
	"github.com/san-lab/sx402/evmbinding"
	"github.com/san-lab/sx402/signing"
)

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

type PermitNonceQuery struct {
	Network string `form:"network" binding:"required"`
	Asset   string `form:"asset" binding:"required"`
	Owner   string `form:"owner" binding:"required"`
}

var clients = evmbinding.InitClients()

func permitNonceHandler(c *gin.Context) {
	var query PermitNonceQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nonce, err := evmbinding.PermitNonce(query.Network, query.Asset, query.Owner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("failed to get nonce: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"network": query.Network,
		"asset":   query.Asset,
		"owner":   query.Owner,
		"nonce":   nonce.String(),
	})
}
