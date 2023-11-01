package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/goccy/go-json"
	"net/http"
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

		if err := h.ctrl.AddFriend(c.Request.Context(), input.Emails); err != nil {
			CustomError(c, err)
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Add friend successfully!"})
		}
	}
}
