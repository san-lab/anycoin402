package all712

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var eip712DomainTypeHash = crypto.Keccak256Hash([]byte("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"))
var transferTypeHash = crypto.Keccak256Hash([]byte("TransferWithAuthorization(address from,address to,uint256 value,uint256 validAfter,uint256 validBefore,bytes32 nonce)"))

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

	// --- Domain separator ---
	domainArgs := abi.Arguments{
		{Type: mustNewType("bytes32")},
		{Type: mustNewType("bytes32")},
		{Type: mustNewType("bytes32")},
		{Type: mustNewType("uint256")},
		{Type: mustNewType("address")},
	}
	domainPacked, err := domainArgs.Pack(
		eip712DomainTypeHash,
		crypto.Keccak256Hash([]byte(tokenName)),
		crypto.Keccak256Hash([]byte(tokenVersion)),
		chainID,
		tokenAddress,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack domain separator: %v", err)
	}
	domainSeparator := crypto.Keccak256Hash(domainPacked)
	log.Println("Domain Separator:", domainSeparator)
	// --- Struct hash for TransferWithAuthorization ---
	transferArgs := abi.Arguments{
		{Type: mustNewType("bytes32")},
		{Type: mustNewType("address")},
		{Type: mustNewType("address")},
		{Type: mustNewType("uint256")},
		{Type: mustNewType("uint256")},
		{Type: mustNewType("uint256")},
		{Type: mustNewType("bytes32")},
	}

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

	snonce, err := hex.DecodeString(auth.Nonce)
	if err != nil {
		return nil, errors.New("Wrong nonce: " + auth.Nonce)
	}

	var nonce [32]byte
	copy(nonce[:], snonce)

	transferPacked, err := transferArgs.Pack(
		transferTypeHash,
		from,
		to,
		value,
		validAfter,
		validBefore,
		nonce,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack struct hash: %v", err)
	}
	structHash := crypto.Keccak256Hash(transferPacked)

	// --- Final EIP-712 digest ---
	prefix := []byte{0x19, 0x01}
	eip712Bytes := append(prefix, append(domainSeparator.Bytes(), structHash.Bytes()...)...)
	digest := crypto.Keccak256Hash(eip712Bytes)

	// --- Sign ---
	signature, err := crypto.Sign(digest.Bytes(), privateKey)
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

	nonce := crypto.Keccak256([]byte{42})
	ppld.Payload.Authorization.Nonce = hex.EncodeToString(nonce)
	auth := Authorization(*ppld.Payload.Authorization)
	bts, err := SignERC3009Authorization(&auth, from_key, chainID, extra["name"], extra["version"], asset)
	if err != nil {
		return nil, fmt.Errorf("Error signing: %w", err)
	}

	ppld.Payload.Signature = hex.EncodeToString(bts)

	return ppld, nil
}

var ChainIDs = map[string]*big.Int{"base-sepolia": big.NewInt(84532)}
