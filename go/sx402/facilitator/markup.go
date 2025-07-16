package facilitator

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/evmbinding"
	"github.com/san-lab/sx402/schemes"
)

type MarkupQuery struct {
	Network string `form:"network" binding:"required"`
	Scheme  string `form:"scheme" binding:"required"`
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

	markup, err := evmbinding.GetMarkup(query.Network, scheme.Asset, keyfile.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("failed to get nonce: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"network": query.Network,
		"scheme":  query.Scheme,
		"markup":  markup.String(),
	})
}
