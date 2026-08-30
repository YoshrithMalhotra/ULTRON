package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"ultron/internal/assistant"
	"ultron/internal/llm"
	"ultron/internal/ratelimit"
)

func main() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	apiKey := os.Getenv("GROQ_API_KEY")

	llmClient := llm.NewClient(
		apiKey,
		"https://api.groq.com/openai/v1/chat/completions",
	)

	ultron := &assistant.Assistant{
		LLM: llmClient,
	}

	chatLimiter := ratelimit.New(rate.Every(6*time.Second), 5)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.POST("/chat", chatLimiter.Middleware(), assistant.ChatHandler(ultron))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}

