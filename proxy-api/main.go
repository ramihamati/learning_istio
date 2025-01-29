package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		resp, err := http.Get("http://weatherhost:3001")
		defer resp.Body.Close()

		if err != nil {
			log.Printf("Error making request %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error making request"})
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading response %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error reading response"})
		}

		log.Printf("Response : %s", string(body))
		c.JSON(http.StatusOK, gin.H{"response": string(body)})
	})

	println("Proxy Api Listening on port 8002")
	r.Run(":3002")
}
