package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	ServerError       = errors.New("Internal server error")
	UserAlreadyExist  = errors.New("User Email already exist")
	UserNotFound      = errors.New("User Email not found")
	FriendshipExisted = errors.New("Friendship already added")
	AlreadySubscribed = errors.New("Requester already subscribed to this Target")
	AlreadyBlocked    = errors.New("Requester already blocked this Target")
	BlockedUser       = errors.New("You Blocked This User")
)

func CustomError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, UserAlreadyExist):
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
	case errors.Is(err, UserNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "Email not found"})
	case errors.Is(err, FriendshipExisted):
		c.JSON(http.StatusConflict, gin.H{"error": "Friendship already added"})
	case errors.Is(err, AlreadySubscribed):
		c.JSON(http.StatusConflict, gin.H{"error": "Requester already subscribed to this Target"})
	case errors.Is(err, AlreadyBlocked):
		c.JSON(http.StatusConflict, gin.H{"error": "Requester already blocked this Target"})
	case errors.Is(err, BlockedUser):
		c.JSON(http.StatusBadRequest, gin.H{"error": "You already blocked this user "})
	case errors.Is(err, ServerError):
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}
}
