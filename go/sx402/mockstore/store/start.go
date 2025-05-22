package store

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var StorePrefix = "mockstore"

func Start(router *gin.Engine) {
	log.Println("starting the demo store")
	router.LoadHTMLGlob(StorePrefix + "/templates/*html")
	store := router.Group(StorePrefix)
	// Index page
	store.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	store.GET("/ab", func(c *gin.Context) {
		c.HTML(http.StatusOK, "indexAB.html", nil)
	})

	// Protected resources
	store.GET("/resources", X402Middleware, ResourceHandler)

}
