package evmbinding

import (
	"context"
	"sync"

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
