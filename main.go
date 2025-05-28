package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// a simple gin server example with an potential vulnerability in older versions
	r := gin.Default()

	// main endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
			"time":    time.Now().Format(time.RFC3339),
		})
	})

	// file endpoint
	r.GET("/file/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		c.JSON(http.StatusOK, gin.H{
			"requested_file": filename,
			"message":        "This endpoint demonstrates potential vulnerability in older Gin versions",
		})
	})

	// health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	// run the server
	r.Run(":8080")
}
