package facilitator

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/mockstore/store"
)

func Start(withStore bool) {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4021"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Define the endpoint
	router.GET("/facilitator/*action", func(c *gin.Context) {
		c.Writer.WriteString("You probably want to use POST method for your action: " + c.Param("action"))
	})

	withEnvelope := router.Group("/facilitator", RequestLogger(), ParseEnvelope, SetupClient)
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

	if withStore {
		store.Start(router)
	}
	// Start the server
	router.Run(":3010")
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log Method, URL and Headers
		log.Printf("Incoming request: %s %s", c.Request.Method, c.Request.URL)
		for k, v := range c.Request.Header {
			log.Printf("Header: %s = %v", k, v)
		}

		// Read and log Body
		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			log.Printf("Body: %s", string(bodyBytes))

			// Restore the io.ReadCloser to its original state
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		c.Next()
	}
}
