package handler

import (
	"assignment/internal/controller"
	"net/http"

	"github.com/friendsofgo/errors"
	"github.com/gin-gonic/gin"
)

var (
	EmailNotFound     = "Email not found"
	UserAlreadyExist  = "User already exists"
	FriendshipExisted = "Friendship already added"
	ServerError       = "Internal Server Error"
	AlreadySubscribed = "Requester already subscribed to this Target"
	AlreadyBlocked    = "Requester already blocked this Target"
	UnidentifiedError = "Something went wrong"
)

func CustomError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, controller.UserAlreadyExist):
		c.JSON(http.StatusConflict, gin.H{"error": UserAlreadyExist})
	case errors.Is(err, controller.UserNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": EmailNotFound})
	case errors.Is(err, controller.FriendshipExisted):
		c.JSON(http.StatusConflict, gin.H{"error": FriendshipExisted})
	case errors.Is(err, controller.ServerError):
		c.JSON(http.StatusInternalServerError, gin.H{"error": ServerError})
	case errors.Is(err, controller.AlreadySubscribed):
		c.JSON(http.StatusConflict, gin.H{"error": AlreadySubscribed})
	case errors.Is(err, controller.AlreadyBlocked):
		c.JSON(http.StatusConflict, gin.H{"error": AlreadyBlocked})
	case errors.Is(err, controller.BlockedUser):
		c.JSON(http.StatusBadRequest, gin.H{"error": AlreadyBlocked})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": UnidentifiedError})
	}
}
