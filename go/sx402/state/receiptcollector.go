package state

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/san-lab/sx402/evmbinding"
)

type OmniHash struct {
	Hash    common.Hash
	Network string
}

type PendingReceipt struct {
	SubmittedAt  time.Time
	TimeToSettle time.Duration
	Receipt      *types.Receipt
}

type ReceiptTracker struct {
	clients  map[string]*ethclient.Client
	receipts sync.Map // map[OmniHash]*pendingReceipt
}

const receiptTimeout = 30 * time.Minute
const pollInterval = 5 * time.Second

func NewReceiptTracker() *ReceiptTracker {
	rt := &ReceiptTracker{
		clients: evmbinding.InitClients(),
	}
	go rt.pollLoop()
	return rt
}

// Submit a tx hash to begin tracking
func (rt *ReceiptTracker) Submit(hash common.Hash, network string) {
	key := OmniHash{Hash: hash, Network: network}
	rt.receipts.Store(key, &PendingReceipt{
		SubmittedAt: time.Now(),
		Receipt:     nil,
	})
	log.Printf("📩 Submitted tx %s on %s", hash.Hex(), network)
}

// Get retrieves a receipt if available
func (rt *ReceiptTracker) Get(hash common.Hash, network string) (*PendingReceipt, bool) {
	val, ok := rt.receipts.Load(OmniHash{Hash: hash, Network: network})
	if !ok {
		return nil, false
	}
	pr := val.(*PendingReceipt)
	if pr.Receipt == nil {
		return nil, false
	}
	return pr, true
}

func (rt *ReceiptTracker) pollLoop() {
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()

		rt.receipts.Range(func(key, value any) bool {
			omni := key.(OmniHash)
			pr := value.(*PendingReceipt)

			// Already received
			if pr.Receipt != nil {
				return true
			}

			// Timeout expired
			if now.Sub(pr.SubmittedAt) > receiptTimeout {
				log.Printf("⏱️ Timeout: %s (%s) exceeded %v, dropping", omni.Hash.Hex(), omni.Network, receiptTimeout)
				rt.receipts.Delete(omni)
				return true
			}

			// Try fetching receipt
			client, ok := rt.clients[omni.Network]
			if !ok {
				log.Printf("⚠️ No client for network: %s", omni.Network)
				return true
			}

			receipt, err := client.TransactionReceipt(context.Background(), omni.Hash)
			if err != nil {
				log.Printf("🔄 Pending: %s (%s)", omni.Hash.Hex(), omni.Network)
				return true
			}

			// Fetch block to get timestamp
			block, err := client.HeaderByNumber(context.Background(), receipt.BlockNumber)
			if err == nil {

				settleTime := time.Unix(int64(block.Time), 0)
				pr.TimeToSettle = settleTime.Sub(pr.SubmittedAt)
			} else {
				log.Println("Failed to get block from ", omni.Network)
				pr.TimeToSettle = time.Since(pr.SubmittedAt)
			}

			pr.Receipt = receipt
			rt.receipts.Store(omni, pr)
			log.Printf("✅ Receipt for %s (%s) stored", omni.Hash.Hex(), omni.Network)
			return true
		})
	}
}
