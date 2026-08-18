package assistant

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChatRequest struct {
	Message string `json:"message"`
}

func ChatHandler(c *gin.Context) {
	var req ChatRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	assistant := &Assistant{}
	response := assistant.Respond(req.Message)

	c.JSON(http.StatusOK, gin.H{
		"response": response,
	})
}
