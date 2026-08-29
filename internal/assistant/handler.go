package assistant

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ultron/internal/llm"
)

type ChatRequest struct {
	Messages []llm.Message `json:"messages"`
}

func ChatHandler(a *Assistant) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ChatRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if len(req.Messages) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "messages is required",
			})
			return
		}

		response, err := a.Respond(req.Messages)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"response": response,
		})
	}
}
