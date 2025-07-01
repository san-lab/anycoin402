package facilitator

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/evmbinding"
)

type PermitNonceQuery struct {
	Network string `form:"network" binding:"required"`
	Asset   string `form:"asset" binding:"required"`
	Owner   string `form:"owner" binding:"required"`
}

var clients = evmbinding.InitClients()

func permitNonceHandler(c *gin.Context) {
	var query PermitNonceQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nonce, err := evmbinding.PermitNonce(query.Network, query.Asset, query.Owner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("failed to get nonce: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"network": query.Network,
		"asset":   query.Asset,
		"owner":   query.Owner,
		"nonce":   nonce.String(),
	})
}
