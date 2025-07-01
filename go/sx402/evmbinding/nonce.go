package evmbinding

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type nonceState struct {
	mu    sync.Mutex
	nonce uint64
}

var nonceMap sync.Map // map[common.Address]*nonceState

func getNonce(ctx context.Context, client *ethclient.Client, address common.Address) (uint64, error) {
	// Load or initialize nonceState for this address
	val, _ := nonceMap.LoadOrStore(address, &nonceState{})
	state := val.(*nonceState)

	state.mu.Lock()
	defer state.mu.Unlock()

	// Fetch the latest pending nonce from the node
	inblock, err := client.PendingNonceAt(ctx, address)
	if err != nil {
		return 0, err
	}

	// Sync in-memory nonce with on-chain nonce if on-chain is higher
	if inblock > state.nonce {
		state.nonce = inblock
	}

	nonceToUse := state.nonce
	state.nonce++

	return nonceToUse, nil
}

func PermitNonce(network, asset, owner string) (*big.Int, error) {
	rpc, ok := GetRPCEndpoint(network)
	if !ok {
		return nil, fmt.Errorf("unsupported network: %s", network)
	}
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to rpc: %w", err)
	}
	defer client.Close()

	tokenAddress := common.HexToAddress(asset)
	ownerAddress := common.HexToAddress(owner)

	// Minimal ABI for nonces function (EIP-2612)

	parsedABI, err := abi.JSON(strings.NewReader(erc20PermitABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	data, err := parsedABI.Pack("nonces", ownerAddress)
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
	var nonce *big.Int
	err = parsedABI.UnpackIntoInterface(&nonce, "nonces", output)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack output: %w", err)
	}

	return nonce, nil

}

const erc20PermitABI = `[{"constant":true,"inputs":[{"name":"owner","type":"address"}],"name":"nonces","outputs":[{"name":"","type":"uint256"}],"type":"function"}]`
