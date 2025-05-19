package evmbinding

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const base_sepolia = "base-sepolia"

var RpcEndpoints = map[string]string{base_sepolia: "https://sepolia.base.org"}

func SendTransaction(client *ethclient.Client, signedTx *types.Transaction) (*common.Hash, error) {

	err := client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, fmt.Errorf("could not send tx: %v", err)
	}
	h := signedTx.Hash()
	return &h, nil
}

func CheckTokenBalance(client *ethclient.Client, tokenAddress, ownerAddress common.Address) (*big.Int, error) {

	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(tokenABI))
	if err != nil {
		return nil, fmt.Errorf("Failed to parse ABI: %v", err)
	}

	// Pack the input (balanceOf(address))
	data, err := parsedABI.Pack("balanceOf", ownerAddress)
	if err != nil {
		log.Fatalf("Failed to pack data: %v", err)
	}

	// Prepare the call message
	msg := ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}

	// Call the contract
	ctx := context.Background()
	result, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		log.Fatalf("Failed to call contract: %v", err)
	}

	// Unpack the result
	var balance *big.Int
	err = parsedABI.UnpackIntoInterface(&balance, "balanceOf", result)

	return balance, err
}

// Returns if the nonce is "known"
func CheckAuthorizationState(client *ethclient.Client, tokenAddress, payer common.Address, nonce [32]byte) (bool, error) {
	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(tokenABI))
	if err != nil {
		return false, fmt.Errorf("Failed to parse ABI: %v", err)
	}

	// Pack the input (balanceOf(address))
	data, err := parsedABI.Pack("authorizationState", payer, nonce)
	if err != nil {
		log.Fatalf("Failed to pack data: %v", err)
	}

	// Prepare the call message
	msg := ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}

	// Call the contract
	ctx := context.Background()
	result, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		log.Fatalf("Failed to call contract: %v", err)
	}

	// Unpack the result

	unpacked, err := parsedABI.Unpack("authorizationState", result)
	known, ok := unpacked[0].(bool)
	if !ok {
		log.Fatalf("Unpacked result is not a bool")
	}

	return known, err
}

func TransferWithAuthorization(
	client *ethclient.Client,
	signer *ecdsa.PrivateKey,
	token, from, to common.Address,
	value, validAfter, validBefore *big.Int,
	nonce, r, s [32]byte,
	v byte,
) (*common.Hash, error) {
	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(trWithAuthABI))
	if err != nil {
		return nil, err
	}

	// Pack the input (balanceOf(address))
	data, err := parsedABI.Pack("transferWithAuthorization", from, to, value, validAfter, validBefore, nonce, v, r, s)
	if err != nil {
		return nil, fmt.Errorf("Failed to pack data: %w", err)
	}

	facilAddress := crypto.PubkeyToAddress(signer.PublicKey)

	// Get nonce for transaction
	fromNonce, err := client.PendingNonceAt(context.Background(), facilAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Estimate gas
	gasLimit := uint64(100000) // or estimate with client.EstimateGas()

	tx := types.NewTransaction(
		fromNonce,
		token,
		big.NewInt(0), // No ETH being sent
		gasLimit,
		gasPrice,
		data,
	)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error recoovering ChainID: %w", err)
	}
	// Sign transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), signer)
	if err != nil {
		return nil, fmt.Errorf("error signing transaction: %w", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, fmt.Errorf("could not send tx: %v", err)
	}
	h := signedTx.Hash()
	return &h, nil

}

const tokenABI = `[
  {
    "constant": true,
    "inputs": [{"name": "_owner", "type": "address"}],
    "name": "balanceOf",
    "outputs": [{"name": "balance", "type": "uint256"}],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "constant": true,
    "inputs": [
      {"name": "authorizer", "type": "address"},
      {"name": "nonce", "type": "bytes32"}
    ],
    "name": "authorizationState",
    "outputs": [{"name": "", "type": "bool"}],
    "stateMutability": "view",
    "type": "function"
  }
]`

const trWithAuthABI = `[{
	"constant": false,
	"inputs": [
	  { "name": "from", "type": "address" },
	  { "name": "to", "type": "address" },
	  { "name": "value", "type": "uint256" },
	  { "name": "validAfter", "type": "uint256" },
	  { "name": "validBefore", "type": "uint256" },
	  { "name": "nonce", "type": "bytes32" },
	  { "name": "v", "type": "uint8" },
	  { "name": "r", "type": "bytes32" },
	  { "name": "s", "type": "bytes32" }
	],
	"name": "transferWithAuthorization",
	"outputs": [],
	"stateMutability": "nonpayable",
	"type": "function"
  }]`
