package facilitator

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"log"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	kms "github.com/proveniencenft/kmsclitool/common"
	"github.com/san-lab/sx402/all712"
	"github.com/san-lab/sx402/evmbinding"
)

var keyfile_name = "0xfAc178B1C359D41e9162A1A6385380de96809048.json"

func InitKeys(password []byte) error {
	kf, err := kms.ReadKeyfile(keyfile_name)
	if err != nil || kf == nil {

		return fmt.Errorf("Error loading keyfile %s: %w", keyfile_name, err)
	}
	err = kf.Decrypt(password)
	if err != nil {
		log.Fatal("Error decrypting keyfile", keyfile_name, err)
	}
	fpk, err = crypto.ToECDSA(kf.Plaintext)
	if err != nil {
		return err
	}
	fmt.Println(crypto.PubkeyToAddress(fpk.PublicKey))

	return nil
}

var fpk *ecdsa.PrivateKey

func SettleHandler(c *gin.Context) {
	if fpk == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Facilitator not properly initialized"})
		c.Abort()
		return
	}

	response := types.SettleResponse{}

	enlp, exists := c.Get("envelope")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Envelope not found"})
		return
	}
	payload := enlp.(all712.Envelope)

	clnt, exists := c.Get("client")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No Client from middleware"})
		return
	}

	client := clnt.(*ethclient.Client)

	response.Network = payload.PaymentPayload.Network

	var from, to, tokenAddress common.Address
	var value, validAfter, validBefore *big.Int

	var nonce, r, s [32]byte
	var v byte

	from = common.HexToAddress(payload.PaymentPayload.Payload.Authorization.From)
	to = common.HexToAddress(payload.PaymentPayload.Payload.Authorization.To)
	tokenAddress = common.HexToAddress(payload.PaymentRequirements.Asset)

	// Convert value
	value, ok := new(big.Int).SetString(payload.PaymentPayload.Payload.Authorization.Value, 10)
	if !ok {
		reason := "error Invalid value format"
		response.ErrorReason = &reason
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Convert validAfter / validBefore
	validAfter, ok = new(big.Int).SetString(payload.PaymentPayload.Payload.Authorization.ValidAfter, 10)
	if !ok {
		reason := "error Invalid ValidAfter format"
		response.ErrorReason = &reason
		c.JSON(http.StatusBadRequest, response)
	}

	validBefore, ok = new(big.Int).SetString(payload.PaymentPayload.Payload.Authorization.ValidBefore, 10)
	if !ok {
		reason := "error Invalid ValidBefore format"
		response.ErrorReason = &reason
		c.JSON(http.StatusBadRequest, response)
	}

	// Convert nonce (hex string to [32]byte)
	nonceBytes, err := hex.DecodeString(strings.TrimPrefix(payload.PaymentPayload.Payload.Authorization.Nonce, "0x"))
	if err != nil || len(nonceBytes) != 32 {
		reason := "error: Invalid nonce"
		response.ErrorReason = &reason
		c.JSON(http.StatusBadRequest, response)
		return
	}
	copy(nonce[:], nonceBytes)

	// Convert r, s (hex strings to []byte)
	sig, err := hex.DecodeString(strings.TrimPrefix(payload.PaymentPayload.Payload.Signature, "0x"))
	if err != nil || len(r) != 32 {
		reason := "error: Invalid signature format"
		response.ErrorReason = &reason
		c.JSON(http.StatusBadRequest, response)
		return
	}
	copy(r[:], sig[:32])
	copy(s[:], sig[32:64])
	v = sig[64]
	if v < 27 {
		v += 27
	}

	h, err := evmbinding.TransferWithAuthorization(client, fpk, tokenAddress, from, to, value, validAfter, validBefore, nonce, r, s, v)

	if err != nil {
		log.Println("error executing settlement", err)
		reason := fmt.Sprintf("Error parsing ABI: %s", err.Error())
		response.ErrorReason = &reason
		c.JSON(http.StatusInternalServerError, response)
		//c.Abort()
		return
	}

	response.Success = true
	response.Transaction = h.Hex()
	payer := from.Hex()
	response.Payer = &payer
	c.JSON(http.StatusOK, response)

}
