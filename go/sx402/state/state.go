package state

import (
	"github.com/ethereum/go-ethereum/common"
)

var receiptCollector = NewReceiptTracker()

func GetReceiptCollector() *ReceiptTracker {
	return receiptCollector
}

func GetPendingReceipt(tx common.Hash, network string) (*PendingReceipt, bool) {
	if receiptCollector == nil {
		return nil, false
	}
	return receiptCollector.Get(tx, network)
}
