package controller

import (
	"assignment/internal/model"
	"assignment/internal/repository"
	"context"
)

type Controller interface {
	AddUsers(context.Context, model.User) error
	AddFriends(context.Context, []string) error
	FriendsList(context.Context, string) (model.FriendshipInfo, error)
	CommonFriends(context.Context, []string) (model.FriendshipInfo, error)
	BlockUsers(context.Context, []string) error
	AddSubscription(context.Context, []string) error
}
type CTRLImplement struct {
	repo repository.Repository
}

func New(repo repository.Repository) Controller {
	return CTRLImplement{repo}
}