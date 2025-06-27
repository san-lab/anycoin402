package all712

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type EVMPayload types.ExactEvmPayload

/*
type Payload struct {
	Signature     string         `json:"signature"`
	Authorization *Authorization `json:"authorization"`
}

type Authorization struct {
	From        string   `json:"from"`
	To          string   `json:"to"`
	Value       string   `json:"value"`
	ValidAfter  *big.Int `json:"validAfter"`
	ValidBefore *big.Int `json:"validBefore"`
	Nonce       string   `json:"nonce"`
}
*/

func VerifyTransferWithAuthorizationSignature(
	signatureHex string,
	auth types.ExactEvmPayloadAuthorization,
	name string,
	version string,
	chainID *big.Int,
	tokenAddress common.Address,
) (bool, common.Address, error) {

	// Hash type: keccak256("TransferWithAuthorization(address from,address to,uint256 value,uint256 validAfter,uint256 validBefore,bytes32 nonce)")

	value, ok := new(big.Int).SetString(auth.Value, 10)
	if !ok {
		return false, common.Address{}, errors.New("Invalid Value")
	}
	after, ok := new(big.Int).SetString(auth.ValidAfter, 10)
	if !ok {
		return false, common.Address{}, errors.New("Invalid After")
	}
	before, ok := new(big.Int).SetString(auth.ValidBefore, 10)
	if !ok {
		return false, common.Address{}, errors.New("Invalid Before")
	}

	nonce_s, err := hex.DecodeString(strings.TrimPrefix(auth.Nonce, "0x"))
	if err != nil {
		return false, common.Address{}, errors.New("Invalid nonce")
	}
	var nonce_h [32]byte
	copy(nonce_h[:], nonce_s)

	digest, err := EIP721Hash(
		common.HexToAddress(auth.From),
		common.HexToAddress(auth.To),
		tokenAddress,
		value,
		after,
		before,
		chainID,
		nonce_h,
		name,
		version,
	)
	if err != nil {
		return false, common.Address{}, err
	}

	// Decode signature
	sig, err := hex.DecodeString(strings.TrimPrefix(signatureHex, "0x"))
	if err != nil {
		return false, common.Address{}, err
	}
	if len(sig) != 65 {
		return false, common.Address{}, fmt.Errorf("invalid signature length")
	}

	// Adjust V if needed
	if sig[64] >= 27 {
		sig[64] -= 27
	}

	pubKey, err := crypto.SigToPub(digest, sig)
	if err != nil {
		return false, common.Address{}, err
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	// Compare recovered address with `from`
	isValid := strings.Compare(strings.ToLower(recoveredAddr.Hex()), strings.ToLower(auth.From)) == 0
	log.Println("Recovered:", recoveredAddr)
	log.Println("From:", auth.From)
	if !isValid {
		return false, recoveredAddr, fmt.Errorf("Recovered address differ: %s expected %s", recoveredAddr, auth.From)
	}

	return true, recoveredAddr, nil
}

func mustNewType(t string) abi.Type {
	typ, err := abi.NewType(t, "", nil)
	if err != nil {
		log.Fatalf("failed to create ABI type %s: %v", t, err)
	}
	return typ
}

type Envelope struct {
	X402Version         int                        `json:"x402Version"`
	PaymentPayload      *PaymentPayload            `json:"paymentPayload"`
	PaymentRequirements *types.PaymentRequirements `json:"paymentRequirements"`
}

type PaymentPayload struct {
	X402Version int             `json:"x402Version"`
	Scheme      string          `json:"scheme"`
	Network     string          `json:"network"`
	Payload     json.RawMessage `json:"payload"`
}
