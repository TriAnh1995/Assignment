package controller

import (
	"assignment/internal/model"
	"assignment/internal/repository"
	"context"
)

type Controller interface {
	AddUsers(context.Context, model.User) error
	AddFriend(context.Context, []string) error
}

type CTRLImplement struct {
	repo repository.Repository
}

func New(repo repository.Repository) Controller {
	return CTRLImplement{repo}
}

func (i CTRLImplement) AddFriend(ctx context.Context, userEmails []string) error {
	return nil
}
