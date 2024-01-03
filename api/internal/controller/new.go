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
}

type CTRLImplement struct {
	repo repository.Repository
}

func New(repo repository.Repository) Controller {
	return CTRLImplement{repo}
}

func (i CTRLImplement) FriendsList(ctx context.Context, userEmail string) (model.FriendshipInfo, error) {
	return model.FriendshipInfo{}, nil
}

func (i CTRLImplement) CommonFriends(ctx context.Context, data []string) (model.FriendshipInfo, error) {
	return model.FriendshipInfo{}, nil
}

