package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FriendsList struct {
	Email string `json:"email"`
}

func (h Handler) FriendsList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input FriendsList
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get your information"})
			return
		}

		if err := input.validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := h.ctrl.FriendsList(c.Request.Context(), input.Email)
		if err != nil {
			CustomError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Success!"})
		c.IndentedJSON(200, result)
	}
}
