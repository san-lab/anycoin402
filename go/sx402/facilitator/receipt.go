package facilitator

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/state"
)

func HandlerReceiptStatus(c *gin.Context) {
	network := c.Query("network")
	tx := c.Query("tx")

	if network == "" || tx == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing required query params: network, tx",
		})
		return
	}

	hash := common.HexToHash(tx)

	pr, ok := state.GetPendingReceipt(hash, network) //rt.receipts.Load(key)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"status": "not_found",
		})
		return
	}

	await := time.Since(pr.SubmittedAt).Seconds()

	if pr.Receipt == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":     "pending",
			"await_time": await,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "found",
		"settle_time": fmt.Sprintf("%v sec", pr.TimeToSettle.Seconds()),
		"receipt":     pr.Receipt, // Gin uses JSON tags from the receipt struct
	})
}

type RecDisplayData struct {
	Network    string
	Tx         string
	Submitted  string
	Error      string
	Status     string
	SettleTime string
	Receipt    string
}

func prettyReceiptPage(c *gin.Context) {

	network := c.Query("network")
	tx := c.Query("tx")

	recdata := new(RecDisplayData)
	recdata.Network = network
	recdata.Tx = tx

	if network == "" || tx == "" {

		recdata.Error = "Missing required query params: network, tx"

	} else {
		hash := common.HexToHash(tx)

		pr, ok := state.GetPendingReceipt(hash, network) //rt.receipts.Load(key)
		if !ok {
			recdata.Status = "not_found"

		} else {
			rec, _ := json.MarshalIndent(pr.Receipt, " ", " ")
			recdata.Submitted = fmt.Sprintf("%s", pr.SubmittedAt.Unix())
			recdata.Status = "found"
			recdata.SettleTime = fmt.Sprintf("%v sec", pr.TimeToSettle.Seconds())
			recdata.Receipt = string(rec)
		}

	}

	c.HTML(http.StatusOK, "receipt.html", recdata)
}
