package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Friends struct {
	Emails []string `json:"friends"`
}

func (h Handler) AddFriend() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input Friends
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get your information"})
			return
		}

		if err := input.validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := h.ctrl.AddFriends(c.Request.Context(), input.Emails); err != nil {
			CustomError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Add friend successfully!"})

	}
}
