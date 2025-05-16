package evmbinding

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const base_sepolia = "base-sepolia"

var rpcendpoints = map[string]string{base_sepolia: "https://sepolia.base.org"}
var baseSepoliaChainId = 84532

func SendTransaction(network string, signedTx *types.Transaction) (*common.Hash, error) {
	url, ok := rpcendpoints[network]
	if !ok {
		return nil, errors.New("Unknown network: " + network)
	}

	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("could not connect to rpc: %v", err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, fmt.Errorf("could not send tx: %v", err)
	}
	h := signedTx.Hash()
	return &h, nil
}

func CheckTokenBalance(network string, tokenAddress, ownerAddress common.Address) (*big.Int, error) {
	url, ok := rpcendpoints[network]
	if !ok {
		return nil, errors.New("Unknown network: " + network)
	}

	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("could not connect to rpc: %v", err)
	}

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
