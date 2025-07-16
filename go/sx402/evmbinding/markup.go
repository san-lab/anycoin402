package evmbinding

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const markupABI = `[{
  "constant": true,
  "inputs": [
    {
      "name": "",
      "type": "address"
    }
  ],
  "name": "markups",
  "outputs": [
    {
      "name": "",
      "type": "uint256"
    }
  ],
  "payable": false,
  "stateMutability": "view",
  "type": "function"
}]`

func GetMarkup(network, asset, facilitator string) (*big.Int, error) {
	client, err := GetClientByNetwork(network)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to rpc: %w", err)
	}
	defer client.Close()

	tokenAddress := common.HexToAddress(asset)
	facAddress := common.HexToAddress(facilitator)

	// Minimal ABI for nonces function (EIP-2612)

	parsedABI, err := abi.JSON(strings.NewReader(markupABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	data, err := parsedABI.Pack("markups", facAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to pack input data: %w", err)
	}

	callMsg := ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}

	ctx := context.Background()
	output, err := client.CallContract(ctx, callMsg, nil)
	if err != nil {
		return nil, fmt.Errorf("contract call failed: %w", err)
	}

	// unpack output (uint256 nonce)
	var markup *big.Int
	err = parsedABI.UnpackIntoInterface(&markup, "markups", output)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack output: %w", err)
	}

	return markup, nil

}
