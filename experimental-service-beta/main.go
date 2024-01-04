package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router
	router := gin.Default()
	router.Use(corsMiddleware())
	// Define a route and its handler
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, this is a Gin server!")
	})

	router.GET("/connect-to-beta", func(c *gin.Context) {
		client := &http.Client{
			Transport: &http.Transport{},
		}
		// URL to send the GET request to
		url := "http://experimental-service-delta:9003/connect-to-delta" 


		// Create a GET request
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal("Error creating request:", err)
			return
		}
		req.Close = true

		// Send the GET request
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("Error sending request:", err)
			return
		}
		defer resp.Body.Close()

		// Check the response status code
		fmt.Println("Response Status:", resp.Status)
	})

	router.GET("/connect-to-delta", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, this is a Gin server!")
	})

	router.GET("/connect-to-omega", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, this is a Gin server!")
	})
	// Run the server on port 8080
	router.Run(":9002")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400") // 24 hrs

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}