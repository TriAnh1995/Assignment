package controller

import (
	"assignment/internal/middleware"
	"context"

	"gorm.io/gorm"
)

func (i CTRLImplement) AddFriend(ctx context.Context, userEmails []string) error {
	tx, err := middleware.GormConnection()
	if err != nil {
		return middleware.ServerError
	}
	return tx.Transaction(func(db *gorm.DB) error {
		if err = i.checkEmail(ctx, userEmails[0]); err != nil {
			return err
		}
		if err = i.checkEmail(ctx, userEmails[1]); err != nil {
			return err
		}

		checkFriendshipExist, err := i.repo.CheckFriendship(ctx, userEmails)
		if err != nil {
			return middleware.ServerError
		}
		if checkFriendshipExist {
			return middleware.FriendshipExisted
		}

		if err = i.repo.AddFriendship(ctx, userEmails); err != nil {
			return middleware.ServerError
		}
		return nil
	})
}
