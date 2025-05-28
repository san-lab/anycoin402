package state

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var receiptCollector = NewReceiptTracker()

func GetReceiptCollector() *ReceiptTracker {
	return receiptCollector
}

func GetReceipt(tx common.Hash, network string) (*types.Receipt, bool) {
	if receiptCollector == nil {
		return nil, false
	}
	return receiptCollector.Get(tx, network)
}
