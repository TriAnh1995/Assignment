package controller

import (
	"assignment/internal/middleware"
	"assignment/internal/model"
	"context"
)

// AddUsers check business logic before add user input to db
func (i CTRLImplement) AddUsers(ctx context.Context, user model.User) error {
	userExists, err := i.repo.CheckUserByEmail(ctx, user.Email)
	if err != nil {
		return middleware.ServerError
	}
	if userExists {
		return middleware.UserAlreadyExist
	}
	if err = i.repo.AddUser(ctx, user); err != nil {
		return middleware.ServerError
	}
	return nil

}
