package handler

import (
	"assignment/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommonFriends struct {
	Emails []string `json:"friends"`
}

func (h Handler) CommonFriends() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input CommonFriends
		input.Emails = []string{c.Query("email1"), c.Query("email2")}

		if err := input.validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		common, err := h.ctrl.CommonFriends(c.Request.Context(), input.Emails)
		if err != nil {
			middleware.CustomError(c, err)
			return
		}
		c.IndentedJSON(200, common)
	}
}
