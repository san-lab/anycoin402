package facilitator

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/evmbinding"
	"github.com/san-lab/sx402/schemes"
)

type MarkupQuery struct {
	Network string  `form:"network" binding:"required"`
	Scheme  string  `form:"scheme" binding:"required"`
	DstEid  *uint32 `form:"dstEid"` // Optional parameter
}

func getMarkup(c *gin.Context) {
	var query MarkupQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	scheme, err := schemes.GetScheme(query.Scheme, query.Network)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "scheme/network pair not found"})
		return
	}
	markup := new(big.Int)

	if query.DstEid == nil {
		markup, err = evmbinding.GetMarkup(query.Network, scheme.Asset, keyfile.Address)

	} else {
		markup, err = evmbinding.GetDetailedMarkup(query.Network, scheme.Asset, *query.DstEid, keyfile.Address)

	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("failed to get the markup: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"network": query.Network,
		"scheme":  query.Scheme,
		"dstEid":  markup.String(),
		"markup":  markup.String(),
	})
}
