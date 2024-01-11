package controller

import (
	"assignment/internal/middleware"
	"context"
)

func (i CTRLImplement) checkEmail(ctx context.Context, userEmail string) error {
	checkEmailExist, err := i.repo.CheckUserByEmail(ctx, userEmail)
	if err != nil {
		return middleware.ServerError
	}
	if !checkEmailExist {
		return middleware.UserNotFound
	}
	return nil
}
