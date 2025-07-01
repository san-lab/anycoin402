package all712

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var eip712DomainTypeHash = crypto.Keccak256Hash([]byte("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"))
var transferTypeHash = crypto.Keccak256Hash([]byte("TransferWithAuthorization(address from,address to,uint256 value,uint256 validAfter,uint256 validBefore,bytes32 nonce)"))
var permitTypeHash = crypto.Keccak256Hash([]byte("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)"))

func EIP712TransferHash(from, to, tokenAddress common.Address, value, after, before, chainID *big.Int, nonce [32]byte, name, version string) ([]byte, error) {

	// Encode struct hash
	arguments := abi.Arguments{
		{Type: mustNewType("address")},
		{Type: mustNewType("address")},
		{Type: mustNewType("uint256")},
		{Type: mustNewType("uint256")},
		{Type: mustNewType("uint256")},
		{Type: mustNewType("bytes32")},
	}

	packed, err := arguments.Pack(
		from,
		to,
		value,
		after,
		before,
		nonce,
	)
	if err != nil {
		return nil, err
	}

	structHash := crypto.Keccak256Hash(append(transferTypeHash.Bytes(), packed...))
	//structHash := crypto.Keccak256Hash(packed)
	// EIP-712 domain separator
	domainSeparator := MakeDomainSeparator(name, version, chainID, tokenAddress)

	// Final digest (EIP-191)
	digestBytes := crypto.Keccak256(
		[]byte("\x19\x01"),
		domainSeparator.Bytes(),
		structHash.Bytes(),
	)
	return digestBytes, nil
}

func EIP712PermitHash(owner, spender, tokenAddress common.Address,
	value, deadline, chainID, nonce *big.Int,
	name, version string,
) ([]byte, error) {

	// Encode struct hash
	arguments := abi.Arguments{
		{Type: mustNewType("address")}, //owner
		{Type: mustNewType("address")}, //spender
		{Type: mustNewType("uint256")}, //value
		{Type: mustNewType("uint256")}, //nonce
		{Type: mustNewType("uint256")}, //deadline
	}

	packed, err := arguments.Pack(
		owner,
		spender,
		value,
		nonce,
		deadline,
	)
	if err != nil {
		return nil, err
	}

	structHash := crypto.Keccak256Hash(append(permitTypeHash.Bytes(), packed...))
	//structHash := crypto.Keccak256Hash(packed)
	// EIP-712 domain separator
	domainSeparator := MakeDomainSeparator(name, version, chainID, tokenAddress)

	// Final digest (EIP-191)
	digestBytes := crypto.Keccak256(
		[]byte("\x19\x01"),
		domainSeparator.Bytes(),
		structHash.Bytes(),
	)
	return digestBytes, nil
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
