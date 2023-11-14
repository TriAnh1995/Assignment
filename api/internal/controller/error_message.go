package controller

import "github.com/friendsofgo/errors"

var (
	ServerError       = errors.New("Internal server error")
	UserAlreadyExist  = errors.New("User Email already exist")
	UserNotFound      = errors.New("User Email not found")
	FriendshipExisted = errors.New("Friendship already added")
	NoFriendsInCommon = errors.New("No Friends in Common")
)
