package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/mockstore/store"
)

func main() {
	r := gin.Default()

	// Load HTML templates
	r.LoadHTMLGlob("templates/*.html")

	// Index page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/ab", func(c *gin.Context) {
		c.HTML(http.StatusOK, "indexAB.html", nil)
	})

	// Protected resources
	r.GET("/resources", store.X402Middleware, resourceHandler)

	r.Run(":4021")
}

func resourceHandler(c *gin.Context) {
	resid := c.Query("RESID")
	c.JSON(http.StatusOK, gin.H{"message": "Access granted to resource " + resid})
}
