package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a router
	router := gin.Default()

	// Health check route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "API is running 🚀",
		})
	})

	// Start server on port 8080
	router.Run(":8080")
}
