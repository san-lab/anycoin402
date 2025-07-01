package store

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func permitNonceProxyHandler(c *gin.Context) {
	network := c.Query("network")
	asset := c.Query("asset")
	owner := c.Query("owner")

	if network == "" || asset == "" || owner == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required query parameters"})
		return
	}

	// Build facilitator URL with query params
	facilitatorURL, err := url.Parse(fmt.Sprintf("%s/permitnonce", facilitatorURI))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid facilitator base URL"})
		return
	}

	q := facilitatorURL.Query()
	q.Set("network", network)
	q.Set("asset", asset)
	q.Set("owner", owner)
	facilitatorURL.RawQuery = q.Encode()

	// Forward the GET request to the facilitator
	resp, err := http.Get(facilitatorURL.String())
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to contact facilitator"})
		return
	}
	defer resp.Body.Close()

	// Copy status code from facilitator response
	c.Status(resp.StatusCode)

	// Copy headers from facilitator response (optional, or filter as needed)
	for key, values := range resp.Header {
		for _, value := range values {
			c.Writer.Header().Add(key, value)
		}
	}

	// Copy body from facilitator response
	if _, err := io.Copy(c.Writer, resp.Body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read facilitator response"})
		return
	}
}
