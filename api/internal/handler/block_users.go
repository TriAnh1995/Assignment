package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Block struct {
	Requester string `json:"requester"`
	Target    string `json:"target"`
}

func (h Handler) BlockUsers() gin.HandlerFunc {

	return func(c *gin.Context) {
		var input Block
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get your information"})
			return
		}

		if err := input.validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var emails = []string{input.Requester, input.Target}
		if err := h.ctrl.BlockUsers(c.Request.Context(), emails); err != nil {
			CustomError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Blocked successfully!"})
	}

}
