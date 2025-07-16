package signing

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/san-lab/sx402/all712"
	"github.com/san-lab/sx402/evmbinding"
)

var endOfMay = big.NewInt(1748735999)
var may13th = big.NewInt(1747130688)

var BaseSepoliaUSDCAddress = common.HexToAddress("0x036CbD53842c5426634e7929541eC2318f3dCF7e")
var BaseSepoliaEURSAddress = common.HexToAddress("0x6Ac14e603A2742fB919248D66c8ecB05D8Aec1e9")

func SignERC3009Authorization(
	auth *types.ExactEvmPayloadAuthorization,
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

	digest, err := all712.EIP712TransferHash(from, to, tokenAddress, value, validAfter, validBefore, chainID, nonce, tokenName, tokenVersion)

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
	ppld.Payload.Authorization.ValidAfter = fmt.Sprint(time.Now().Add(10 * time.Minute).Unix())
	ppld.Payload.Authorization.ValidBefore = fmt.Sprint(time.Now().Add(24 * time.Hour).Unix())
	ppld.Payload.Authorization.Value = paymentReqs.MaxAmountRequired

	extra := map[string]string{}
	err := json.Unmarshal(*paymentReqs.Extra, &extra)
	if err != nil {
		return nil, err
	}

	asset := common.HexToAddress(paymentReqs.Asset)

	chainID, ok := evmbinding.ChainIDs[ppld.Network]
	if !ok {
		return nil, errors.New("Unknown network: " + ppld.Network)
	}
	timestring := time.Now().GoString()
	nonce := crypto.Keccak256([]byte(timestring))
	ppld.Payload.Authorization.Nonce = "0x" + hex.EncodeToString(nonce)
	auth := ppld.Payload.Authorization
	bts, err := SignERC3009Authorization(auth, from_key, chainID, extra["name"], extra["version"], asset)
	if err != nil {
		return nil, fmt.Errorf("Error signing: %w", err)
	}

	ppld.Payload.Signature = hex.EncodeToString(bts)

	return ppld, nil
}

func SignEIP2612Permit(permit *all712.Permit, privateKey *ecdsa.PrivateKey) ([]byte, error) {

	digest, err := permit.Digest()
	if err != nil {
		return digest, err
	}
	// --- Sign ---
	signature, err := crypto.Sign(digest, privateKey)

	signature[64] += 27
	if err != nil {
		return nil, fmt.Errorf("failed to sign digest: %v", err)
	}

	return signature, nil
}

func VerifyPermitSignature(permit *all712.Permit) (recovered common.Address, err error) {
	if len(permit.Signature) != 132 {
		err = fmt.Errorf("wrong signature length: %v", len(permit.Signature))
		return
	}
	digest, err := permit.Digest()
	if err != nil {
		err = fmt.Errorf("errr hashing Parmit: %w", err)
		return
	}
	sig, err := hex.DecodeString(permit.Signature[2:])
	if err != nil {
		err = fmt.Errorf("error decoding permit signature: %w", err)
		return
	}
	if sig[64] >= 27 {
		sig[64] -= 27
	}

	pub, err := crypto.SigToPub(digest, sig)
	if err != nil {
		err = fmt.Errorf("error recovering address: %w", err)
		return
	}

	recovered = crypto.PubkeyToAddress(*pub)
	if recovered.Cmp(permit.Message.Owner) != 0 {
		err = fmt.Errorf("recoverd key does not match the Owner")
	}
	log.Printf("permit signature: 0x%x", permit.Signature)
	return
}

func VerifyTransferWithAuthorizationSignature(
	signatureHex string,
	auth types.ExactEvmPayloadAuthorization,
	name string,
	version string,
	chainID *big.Int,
	tokenAddress common.Address,
) (recovered common.Address, nonce [32]byte, signature []byte, err error) {

	// Hash type: keccak256("TransferWithAuthorization(address from,address to,uint256 value,uint256 validAfter,uint256 validBefore,bytes32 nonce)")

	value, ok := new(big.Int).SetString(auth.Value, 10)
	if !ok {
		err = errors.New("Invalid Value")
		return
	}
	after, ok := new(big.Int).SetString(auth.ValidAfter, 10)
	if !ok {
		err = errors.New("Invalid After")
		return
	}
	before, ok := new(big.Int).SetString(auth.ValidBefore, 10)
	if !ok {
		err = errors.New("Invalid Before")
		return
	}

	nonce_s, err := hex.DecodeString(strings.TrimPrefix(auth.Nonce, "0x"))
	if err != nil {
		err = errors.New("Invalid nonce")
		return
	}

	copy(nonce[:], nonce_s)

	digest, err := all712.EIP712TransferHash(
		common.HexToAddress(auth.From),
		common.HexToAddress(auth.To),
		tokenAddress,
		value,
		after,
		before,
		chainID,
		nonce,
		name,
		version,
	)
	if err != nil {
		return
	}

	// Decode signature
	signature, err = hex.DecodeString(strings.TrimPrefix(signatureHex, "0x"))
	if err != nil {
		err = fmt.Errorf("error decoding signature: %w", err)
		return
	}
	if len(signature) != 65 {
		err = fmt.Errorf("invalid signature length: %v", len(signature))
		return
	}

	adjustedSignature := make([]byte, 65)
	copy(adjustedSignature, signature)

	// Adjust V if needed
	if adjustedSignature[64] >= 27 {
		adjustedSignature[64] -= 27
	}

	pubKey, err := crypto.SigToPub(digest, adjustedSignature)
	if err != nil {
		return
	}
	recovered = crypto.PubkeyToAddress(*pubKey)

	// Compare recovered address with `from`
	isValid := strings.Compare(strings.ToLower(recovered.Hex()), strings.ToLower(auth.From)) == 0
	log.Println("Recovered:", recovered)
	log.Println("From:", auth.From)
	if !isValid {
		err = fmt.Errorf("Recovered address differ: %s expected %s", recovered, auth.From)
		return
	}

	return
}
