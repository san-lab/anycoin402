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
	"github.com/san-lab/sx402/schemas"
)

func Start(withStore bool, facilitatorPassword []byte) {
	err := InitKeys(facilitatorPassword)
	if err != nil {
		log.Fatal("error initializig keys:", err)
		return
	}
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
	router.GET("facilitator/supported", getSupported)
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

func getSupported(c *gin.Context) {
	supported := []schemas.Scheme{}
	for scheme := range schemas.Assets {
		supported = append(supported, scheme)
	}
	c.JSON(http.StatusOK, gin.H{
		"kinds": supported,
	})
}
