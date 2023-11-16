package controller

import (
	"assignment/internal/model"
	"assignment/internal/repository"
	"context"

	"gorm.io/gorm"
)

type Controller interface {
	AddUsers(context.Context, model.User) error
	AddFriend(context.Context, []string) error
}

type CTRLImplement struct {
	repo   repository.Repository
	gormDB *gorm.DB
}

func New(repo repository.Repository, gormDB *gorm.DB) Controller {
	return CTRLImplement{
		repo:   repo,
		gormDB: gormDB,
	}
}
