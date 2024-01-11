package middleware

import "github.com/friendsofgo/errors"

var (
	ServerError       = errors.New("Internal server error")
	UserAlreadyExist  = errors.New("User Email already exist")
	UserNotFound      = errors.New("User Email not found")
	FriendshipExisted = errors.New("Friendship already added")
	AlreadySubscribed = errors.New("Requester already subscribed to this Target")
	AlreadyBlocked    = errors.New("Requester already blocked this Target")
	BlockedUser       = errors.New("You Blocked This User")
	UserNotUpdate     = errors.New("This User not have any update")
)
