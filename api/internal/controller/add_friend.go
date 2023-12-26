package controller

import (
	"context"

	"gorm.io/gorm"
)

func (i CTRLImplement) AddFriend(ctx context.Context, userEmails []string) error {
	tx, err := GormConnection()
	if err != nil {
		return ServerError
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
			return ServerError
		}
		if checkFriendshipExist {
			return FriendshipExisted
		}

		if err = i.repo.AddFriendship(ctx, userEmails); err != nil {
			return ServerError
		}
		return nil
	})
}
