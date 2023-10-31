package handler

import (
	"assignment/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h Handler) AddUsers() gin.HandlerFunc {

	return func(c *gin.Context) {
		var input User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get your information"})
			return
		}

		if err := input.validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := h.ctrl.AddUsers(c.Request.Context(),
			model.User{
				Name:  input.Name,
				Email: input.Email,
			}); err != nil {
			CustomError(c, err)
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Add user successfully!"})
		}
	}
}
