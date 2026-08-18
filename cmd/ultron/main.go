package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ultron/internal/assistant"
)

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.POST("/chat", assistant.ChatHandler)

	router.Run(":8080")
}
