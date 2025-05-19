package all712

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var endOfMay = big.NewInt(1748735999)
var may13th = big.NewInt(1747130688)

var BaseSepoliaUSDCAddress = common.HexToAddress("0x036CbD53842c5426634e7929541eC2318f3dCF7e")
var BaseSepoliaEURSAddress = common.HexToAddress("0x6Ac14e603A2742fB919248D66c8ecB05D8Aec1e9")

func SignERC3009Authorization(
	auth *Authorization,
	privateKey *ecdsa.PrivateKey,
	chainID *big.Int,
	tokenName string,
	tokenVersion string,
	tokenAddress common.Address,
) ([]byte, error) {

	from := common.HexToAddress(auth.From)
	to := common.HexToAddress(auth.To)
	value, ok := new(big.Int).SetString(auth.Value, 10)
	if !ok {
		return nil, errors.New("error parsing value")
	}

	validAfter, ok := new(big.Int).SetString(auth.ValidAfter, 10)
	if !ok {
		return nil, errors.New("error parsing validAfter")
	}

	validBefore, ok := new(big.Int).SetString(auth.ValidBefore, 10)
	if !ok {
		return nil, errors.New("error parsing validBefore")
	}

	snonce, err := hex.DecodeString(strings.TrimPrefix(auth.Nonce, "0x"))
	if err != nil {
		return nil, errors.New("Wrong nonce: " + auth.Nonce)
	}

	var nonce [32]byte
	copy(nonce[:], snonce)

	digest, err := EIP721Hash(from, to, tokenAddress, value, validAfter, validBefore, chainID, nonce, tokenName, tokenVersion)

	// --- Sign ---
	signature, err := crypto.Sign(digest, privateKey)

	signature[64] += 27
	if err != nil {
		return nil, fmt.Errorf("failed to sign digest: %v", err)
	}

	return signature, nil
}

func AddAuthorizationSignature(paymentReqs *types.PaymentRequirements, from_key *ecdsa.PrivateKey) (*types.PaymentPayload, error) {
	from := crypto.PubkeyToAddress(from_key.PublicKey)

	ppld := new(types.PaymentPayload)
	ppld.Network = paymentReqs.Network
	ppld.Scheme = paymentReqs.Scheme
	ppld.Payload = new(types.ExactEvmPayload)
	ppld.X402Version = 1
	ppld.Payload.Authorization = new(types.ExactEvmPayloadAuthorization)
	ppld.Payload.Authorization.From = from.Hex()
	ppld.Payload.Authorization.To = paymentReqs.PayTo
	ppld.Payload.Authorization.ValidAfter = fmt.Sprint(time.Now().Unix())
	ppld.Payload.Authorization.ValidBefore = fmt.Sprint(time.Now().Add(24 * time.Hour).Unix())
	ppld.Payload.Authorization.Value = paymentReqs.MaxAmountRequired

	extra := map[string]string{}
	err := json.Unmarshal(*paymentReqs.Extra, &extra)
	if err != nil {
		return nil, err
	}

	asset := common.HexToAddress(paymentReqs.Asset)

	chainID, ok := ChainIDs[ppld.Network]
	if !ok {
		return nil, errors.New("Unknown network: " + ppld.Network)
	}
	timestring := time.ANSIC
	nonce := crypto.Keccak256([]byte(timestring))
	ppld.Payload.Authorization.Nonce = "0x" + hex.EncodeToString(nonce)
	auth := Authorization(*ppld.Payload.Authorization)
	bts, err := SignERC3009Authorization(&auth, from_key, chainID, extra["name"], extra["version"], asset)
	if err != nil {
		return nil, fmt.Errorf("Error signing: %w", err)
	}

	ppld.Payload.Signature = hex.EncodeToString(bts)

	return ppld, nil
}

var ChainIDs = map[string]*big.Int{"base-sepolia": big.NewInt(84532)}
