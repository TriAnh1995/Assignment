package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Subscription struct {
	Requester string `json:"requester"`
	Target    string `json:"target"`
}

func (h Handler) Subscription() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input Subscription
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get your information"})
			return
		}

		if err := input.validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		emails := []string{input.Requester, input.Target}
		if err := h.ctrl.Subscription(c.Request.Context(), emails); err != nil {
			CustomError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Subscribe successfully!"})
	}
}
