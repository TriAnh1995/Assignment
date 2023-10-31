package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CustomError(c *gin.Context, err error) {
	switch {
	case strings.Contains(err.Error(), "User Email already exist"):
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
	case strings.Contains(err.Error(), "not found"):
		c.JSON(404, gin.H{"error": "Email not found"})
	case strings.Contains(err.Error(), "Friendship already added"):
		c.JSON(http.StatusConflict, gin.H{"error": "Friendship already added"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
}
