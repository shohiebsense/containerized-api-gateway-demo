package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Message string
}

func main() {
	// Initialize the Gin router
	router := gin.Default()
	router.Use(corsMiddleware())
	// Define a route and its handler
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, this is a Gin server!")
	})

	router.GET("/connect-to-omega", func(c *gin.Context) {
		response := Response {
			Message: "Thank you, you have reached the peak",
		}
		c.JSON(http.StatusOK, response)
	})
	// Run the server on port 8080
	router.Run(":9004")

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