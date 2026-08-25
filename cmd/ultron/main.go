package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"ultron/internal/assistant"
	"ultron/internal/llm"
)

func main() {
	router := gin.Default()

	apiKey := os.Getenv("GROQ_API_KEY")

	llmClient := llm.NewClient(
		apiKey,
		"https://api.groq.com/openai/v1/chat/completions",
	)

	ultron := &assistant.Assistant{
		LLM: llmClient,
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.POST("/chat", assistant.ChatHandler(ultron))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
