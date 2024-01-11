package handler

import (
	"assignment/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func CustomError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, middleware.UserAlreadyExist):
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
	case errors.Is(err, middleware.UserNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "Email not found"})
	case errors.Is(err, middleware.FriendshipExisted):
		c.JSON(http.StatusConflict, gin.H{"error": "Friendship already added"})
	case errors.Is(err, middleware.AlreadySubscribed):
		c.JSON(http.StatusConflict, gin.H{"error": "Requester already subscribed to this Target"})
	case errors.Is(err, middleware.AlreadyBlocked):
		c.JSON(http.StatusConflict, gin.H{"error": "Requester already blocked this Target"})
	case errors.Is(err, middleware.BlockedUser):
		c.JSON(http.StatusBadRequest, gin.H{"error": "You already blocked this user "})
	case errors.Is(err, middleware.ServerError):
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}
}
