package store

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/evmbinding"
)

var StorePrefix = "mockstore"

func Start(router *gin.Engine, tmpl *template.Template) {
	log.Println("starting the demo store")

	template.Must(tmpl.ParseGlob(StorePrefix + "/templates/*html"))
	//router.SetHTMLTemplate(tmpl)
	store := router.Group(StorePrefix)
	// Index page

	chainids := map[string]uint64{}

	for k, v := range evmbinding.ChainIDs {
		chainids[k] = v.Uint64()
	}
	store.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"ChainIDs": chainids})
	})

	store.GET("/ab", func(c *gin.Context) {
		c.HTML(http.StatusOK, "indexAB.html", nil)
	})

	// Protected resources
	store.GET("/resources", X402Middleware, ResourceHandler)

	store.GET("/permitnonce", permitNonceProxyHandler)

}
