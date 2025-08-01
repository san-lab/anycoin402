package evmbinding

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/san-lab/sx402/oftcc"
)

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

	token, err := oftcc.NewOftcc(tokenAddress, client)
	if err != nil {
		return
	}

	// Create a new CallOpts (read-only call)
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}
	markup, err = token.LocalMarkups(callOpts, facAddress)
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

	token, err := oftcc.NewOftcc(tokenAddress, client)
	if err != nil {
		return
	}

	// Create a new CallOpts (read-only call)
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}
	markup, err = token.Markups(callOpts, facAddress, destChainId)
	return

}
