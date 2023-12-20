package handler

import (
	"assignment/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func CustomError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, controller.UserAlreadyExist):
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
	case errors.Is(err, controller.UserNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "Email not found"})
	case errors.Is(err, controller.FriendshipExisted):
		c.JSON(http.StatusConflict, gin.H{"error": "Friendship already added"})
	case errors.Is(err, controller.AlreadySubscribed):
		c.JSON(http.StatusConflict, gin.H{"error": "Requester already subscribed to this Target"})
	case errors.Is(err, controller.AlreadyBlocked):
		c.JSON(http.StatusConflict, gin.H{"error": "Requester already blocked this Target"})
	case errors.Is(err, controller.BlockedUser):
		c.JSON(http.StatusBadRequest, gin.H{"error": "You already blocked this user "})
	case errors.Is(err, controller.ServerError):
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}
}
