package state

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/evmbinding"
)

type OmniHash struct {
	Hash    common.Hash
	Network string
}

type pendingReceipt struct {
	submittedAt time.Time
	receipt     *types.Receipt
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
	rt.receipts.Store(key, &pendingReceipt{
		submittedAt: time.Now(),
		receipt:     nil,
	})
	log.Printf("📩 Submitted tx %s on %s", hash.Hex(), network)
}

// Get retrieves a receipt if available
func (rt *ReceiptTracker) Get(hash common.Hash, network string) (*types.Receipt, bool) {
	val, ok := rt.receipts.Load(OmniHash{Hash: hash, Network: network})
	if !ok {
		return nil, false
	}
	pr := val.(*pendingReceipt)
	if pr.receipt == nil {
		return nil, false
	}
	return pr.receipt, true
}

func (rt *ReceiptTracker) pollLoop() {
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()

		rt.receipts.Range(func(key, value any) bool {
			omni := key.(OmniHash)
			pr := value.(*pendingReceipt)

			// Already received
			if pr.receipt != nil {
				return true
			}

			// Timeout expired
			if now.Sub(pr.submittedAt) > receiptTimeout {
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

			pr.receipt = receipt
			rt.receipts.Store(omni, pr)
			log.Printf("✅ Receipt for %s (%s) stored", omni.Hash.Hex(), omni.Network)
			return true
		})
	}
}

func (rt *ReceiptTracker) HandlerReceiptStatus(c *gin.Context) {
	network := c.Query("network")
	tx := c.Query("tx")

	if network == "" || tx == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing required query params: network, tx",
		})
		return
	}

	hash := common.HexToHash(tx)
	key := OmniHash{Hash: hash, Network: network}

	val, ok := rt.receipts.Load(key)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"status": "not_found",
		})
		return
	}

	pr := val.(*pendingReceipt)
	await := time.Since(pr.submittedAt).Seconds()

	if pr.receipt == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":     "pending",
			"await_time": await,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     "found",
		"await_time": await,
		"receipt":    pr.receipt, // Gin uses JSON tags from the receipt struct
	})
}
