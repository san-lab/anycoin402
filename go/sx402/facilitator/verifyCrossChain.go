package facilitator

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/all712"
	"github.com/san-lab/sx402/evmbinding"
	"github.com/san-lab/sx402/signing"
)

func VerifyCrossChainScheme(c *gin.Context, envelope *all712.Envelope) {

	clnt, exists := c.Get("client")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Client not found"})
		c.Abort()
		return
	}
	client := clnt.(*ethclient.Client)

	response := types.VerifyResponse{}
	response.InvalidReason = new(string)

	// TODO: Use ExtraInfo for additional validation?
	ccmsg, _, err := parseCrossChainMessage(envelope)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Parsing error": err.Error()})
		c.Abort()
		return
	}

	rec, err := signing.VerifyCrossChainAuthSignature(ccmsg)
	if err != nil {
		reason := "Error verifying signature: " + err.Error()
		response.IsValid = false
		response.InvalidReason = &reason
		c.JSON(http.StatusOK, response)
		c.Abort()
		return
	}

	markup, insignificant_err := evmbinding.GetDetailedMarkup(envelope.PaymentPayload.Network,
		envelope.PaymentRequirements.Asset,
		uint32(ccmsg.Authorization.DestinationChain.Uint64()),
		keyfile.Address)
	if insignificant_err != nil {
		log.Println(insignificant_err)
	}

	//The authorzed amount must cover the slippage and the minAmount
	if new(big.Int).Sub(ccmsg.Authorization.Amount, markup).Cmp(ccmsg.Authorization.MinimalAmount) < 0 {
		reason := "minAmount not guaranteed"
		response.IsValid = false
		response.InvalidReason = &reason
		c.JSON(http.StatusOK, response)
	}

	//Reuse the EIP3009 verification for now
	pd := ParsedData{}
	pd.chainID = ccmsg.Domain.ChainID
	pd.nonce = common.HexToHash(ccmsg.Authorization.Nonce)
	pd.Amount = ccmsg.Authorization.Amount
	pd.Asset = ccmsg.Domain.VerifyingContract
	pd.ValidAfter = ccmsg.Authorization.ValidAfter
	pd.ValidBefore = ccmsg.Authorization.ValidBefore
	pd.Payer = rec
	valid, reason := Verify3009OnChainConstraints(client, pd)

	response.IsValid = valid
	response.InvalidReason = &reason
	c.JSON(http.StatusOK, response)
}

func parseCrossChainMessage(envelope *all712.Envelope) (ccmsg *all712.CrossChainTransferMessage, extraInfo *ExtraInfo, err error) {
	ccmsg = new(all712.CrossChainTransferMessage)
	err = json.Unmarshal(envelope.PaymentPayload.Payload, ccmsg)
	if err != nil {
		return
	}

	//This may needa more systematic design
	extraInfo = new(ExtraInfo)
	err = json.Unmarshal(*envelope.PaymentRequirements.Extra, extraInfo)
	return

}
