package controller

import (
	"assignment/internal/model"
	"assignment/internal/repository"
	"context"

	"gorm.io/gorm"
)

type Controller interface {
	AddUsers(context.Context, model.User) error
	AddFriends(context.Context, *gorm.DB, []string) error
}

type CTRLImplement struct {
	repo repository.Repository
}

func New(repo repository.Repository) Controller {
	return CTRLImplement{repo}
}
