package facilitator

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	// Define the endpoint
	router.GET("/facilitator/*action", func(c *gin.Context) {
		c.Writer.WriteString("You probably want to use POST method for your action: " + c.Param("action"))
	})

	withEnvelope := router.Group("/facilitator", ParseEnvelope, SetupClient)
	withEnvelope.POST("/verify", verifyHandler)
	withEnvelope.POST("/settle", SettleHandler)
	router.GET("/", func(c *gin.Context) {
		c.Writer.WriteString("Hello there!")
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Route not found", "path": c.Request.RequestURI,
		})
	})

	// Start the server
	router.Run(":3010")
}
