package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/genproto/googleapis/type/date"
)

type FullUsers struct {
	Title        string    `json:"title"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	PasswordHash string    `json:"password"`
	Gender       string    `json:"gender"`
	Email        string    `json:"email"`
	DateOfBirth  date.Date `json:"date_of_birth"`
	PhoneNumber  string    `json:"phone_number"`
}

func (h Handler) FullUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input FullUsers
		var err error
		if err = c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get your information"})
			return
		}
		if input.PasswordHash, err = HashPassword(input.PasswordHash); err != nil {
			c.JSON(500, gin.H{"error": "Server Error"})
			return
		}
		if err = input.validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
}
