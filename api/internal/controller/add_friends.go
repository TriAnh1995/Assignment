package controller

import (
	"context"

	"gorm.io/gorm"
)

func (i CTRLImplement) AddFriends(ctx context.Context, userEmails []string) error {
	tx, err := ConnectToDatabase()
	if err != nil {
		return ServerError
	}
	return tx.Transaction(func(gormDB *gorm.DB) error {
		// Start Business Logic
		// Check whether ot not both users exists in db
		if err = i.checkEmail(ctx, userEmails[0]); err != nil {
			return err
		}
		if err = i.checkEmail(ctx, userEmails[1]); err != nil {
			return err
		}
		// Check if both user already friends in db
		checkFriendshipExist, err := i.repo.CheckFriendship(ctx, userEmails)
		if err != nil {
			return ServerError
		}
		if checkFriendshipExist {
			return FriendshipExisted
		}
		// Add friendship to db
		if err = i.repo.AddFriendship(ctx, userEmails[0], userEmails[1]); err != nil {
			return ServerError
		}
		return nil
	})
}
