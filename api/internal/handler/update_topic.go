package handler

import (
	"assignment/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateTopic struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

func (h Handler) UpdateTopic() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input UpdateTopic
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get your information"})
			return
		}
		mentionedEmail, err := input.validate()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		updateInfo := model.UpdateInfo{
			Sender: input.Sender,
			Text:   input.Text,
		}
		receivedUpdateList, err := h.ctrl.UpdateTopic(c.Request.Context(), updateInfo)
		if err != nil {
			CustomError(c, err)
			return
		}
		result := append(receivedUpdateList, mentionedEmail)
		c.JSON(http.StatusOK, gin.H{"message": "Success: true"})
		c.IndentedJSON(200, result)
	}
}
