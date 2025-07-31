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

const detailedMarkupABI = `[{
  "inputs": [
    {
      "internalType": "address",
      "name": "",
      "type": "address"
    },
    {
      "internalType": "uint32",
      "name": "",
      "type": "uint32"
    }
  ],
  "name": "markups",
  "outputs": [
    {
      "internalType": "uint256",
      "name": "",
      "type": "uint256"
    }
  ],
  "stateMutability": "view",
  "type": "function"
}]`

func GetMarkup(network, asset, facilitator string) (markup *big.Int, err error) {
	markup = big.NewInt(0)
	client, err := GetClientByNetwork(network)
	if err != nil {
		err = fmt.Errorf("failed to connect to rpc: %w", err)
		return
	}
	defer client.Close()

	tokenAddress := common.HexToAddress(asset)
	facAddress := common.HexToAddress(facilitator)

	// Minimal ABI for nonces function (EIP-2612)

	parsedABI, err := abi.JSON(strings.NewReader(markupABI))
	if err != nil {
		err = fmt.Errorf("failed to parse ABI: %w", err)
		return
	}

	data, err := parsedABI.Pack("markups", facAddress)
	if err != nil {

		return
	}

	callMsg := ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}

	ctx := context.Background()
	output, err := client.CallContract(ctx, callMsg, nil)
	if err != nil {
		err = fmt.Errorf("contract call failed: %w", err)
		return
	}

	// unpack output (uint256 nonce)

	err = parsedABI.UnpackIntoInterface(&markup, "markups", output)
	if err != nil {
		err = fmt.Errorf("failed to unpack output: %w", err)
		return
	}

	return

}

func GetDetailedMarkup(network, asset string, destChainId uint32, facilitator string) (markup *big.Int, err error) {
	markup = big.NewInt(0)
	client, err := GetClientByNetwork(network)
	if err != nil {
		err = fmt.Errorf("failed to connect to rpc: %w", err)
		return
	}
	defer client.Close()

	tokenAddress := common.HexToAddress(asset)
	facAddress := common.HexToAddress(facilitator)

	// Minimal ABI for nonces function (EIP-2612)

	parsedABI, err := abi.JSON(strings.NewReader(detailedMarkupABI))
	if err != nil {
		err = fmt.Errorf("failed to parse ABI: %w", err)
		return
	}

	data, err := parsedABI.Pack("markups", facAddress, destChainId)
	if err != nil {

		return
	}

	callMsg := ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}

	ctx := context.Background()
	output, err := client.CallContract(ctx, callMsg, nil)
	if err != nil {
		err = fmt.Errorf("contract call failed: %w", err)
		return
	}

	// unpack output (uint256 nonce)

	err = parsedABI.UnpackIntoInterface(&markup, "markups", output)
	if err != nil {
		err = fmt.Errorf("failed to unpack output: %w", err)
		return
	}

	return

}
