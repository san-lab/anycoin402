package facilitator

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/mockstore/store"
	"github.com/san-lab/sx402/schemes"
)

var Template = template.New("")

func Start(withStore bool, facilitatorPassword []byte) {
	err := InitKeys(facilitatorPassword)
	if err != nil {
		log.Fatal("error initializig keys:", err)
		return
	}
	router := gin.Default()
	template.Must(Template.ParseGlob("templates/*html"))
	router.SetHTMLTemplate(Template)
	//router.LoadHTMLGlob("templates/*html")
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	// Define the endpoint
	router.GET("facilitator/supported", getSupported)
	router.GET("facilitator/receiptraw", HandlerReceiptStatus)
	router.GET("facilitator/receipt", prettyReceiptPage)
	router.GET("facilitator/permitnonce", permitNonceHandler)
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
		store.Start(router, Template)
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
	supported := []schemes.Scheme{}
	for scheme := range schemes.Assets {
		supported = append(supported, scheme)
	}
	c.JSON(http.StatusOK, gin.H{
		"kinds": supported,
	})
}
