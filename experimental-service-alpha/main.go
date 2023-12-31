package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string
}

func main() {
	router := gin.Default()
	router.Use(corsMiddleware())
	router.GET("/", func(c *gin.Context) {

		response := Response {
			Message: "Ok",
		}
		c.JSON(http.StatusOK, response)
	})

	router.GET("/connect-to-alpha", func(c *gin.Context) {
		client := &http.Client{}

		url := "http://experimental-service-beta:9002/connect-to-beta" 

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal("Error creating request:", err)
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("Error sending request:", err)
			return
		}
		defer resp.Body.Close()

		fmt.Println("Response Status:", resp.Status)
	})

	router.Run(":9001")
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