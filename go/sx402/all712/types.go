package all712

import (
	"encoding/hex"
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
type Authorization types.ExactEvmPayloadAuthorization

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
	typeHash := crypto.Keccak256Hash([]byte("TransferWithAuthorization(address from,address to,uint256 value,uint256 validAfter,uint256 validBefore,bytes32 nonce)"))

	// Encode struct hash
	arguments := abi.Arguments{
		{Type: mustNewType("address")},
		{Type: mustNewType("address")},
		{Type: mustNewType("uint256")},
		{Type: mustNewType("uint256")},
		{Type: mustNewType("uint256")},
		{Type: mustNewType("bytes32")},
	}
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

	nonce_s, err := hex.DecodeString(auth.Nonce[2:])
	if err != nil {
		return false, common.Address{}, errors.New("Invalid nonce")
	}
	var nonce_h [32]byte
	copy(nonce_h[:], nonce_s)

	packed, err := arguments.Pack(
		common.HexToAddress(auth.From),
		common.HexToAddress(auth.To),
		value,
		after,
		before,
		nonce_h,
	)
	if err != nil {
		return false, common.Address{}, err
	}
	structHash := crypto.Keccak256Hash(append(typeHash.Bytes(), packed...))

	// EIP-712 domain separator
	domainSeparator := MakeDomainSeparator(name, version, chainID, tokenAddress)

	// Final digest (EIP-191)
	digestBytes := crypto.Keccak256(
		[]byte("\x19\x01"),
		domainSeparator.Bytes(),
		structHash.Bytes(),
	)
	digest := common.BytesToHash(digestBytes)

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

	pubKey, err := crypto.SigToPub(digest.Bytes(), sig)
	if err != nil {
		return false, common.Address{}, err
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	// Compare recovered address with `from`
	isValid := (recoveredAddr.Hex() == auth.From)
	log.Println("Recovered:", recoveredAddr)
	log.Println("From:", auth.From)

	return isValid, recoveredAddr, nil
}

func MakeDomainSeparator(name, version string, chainID *big.Int, verifyingContract common.Address) common.Hash {
	// keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)")
	typeHash := crypto.Keccak256Hash([]byte("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"))

	nameHash := crypto.Keccak256Hash([]byte(name))
	versionHash := crypto.Keccak256Hash([]byte(version))

	arguments := abi.Arguments{

		{Type: mustNewType("bytes32")},
		{Type: mustNewType("bytes32")},
		{Type: mustNewType("uint256")},
		{Type: mustNewType("address")},
	}

	packed, err := arguments.Pack(
		nameHash,
		versionHash,
		chainID,
		verifyingContract,
	)
	if err != nil {
		log.Fatalf("Domain separator packing failed: %v", err)
	}

	return crypto.Keccak256Hash(append(typeHash.Bytes(), packed...))
}

func mustNewType(t string) abi.Type {
	typ, err := abi.NewType(t, "", nil)
	if err != nil {
		log.Fatalf("failed to create ABI type %s: %v", t, err)
	}
	return typ
}
