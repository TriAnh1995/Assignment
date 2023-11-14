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
	return tx.Transaction(func(*gorm.DB) error {
		// Start Business Logic
		// Check whether ot not both users exists in db with for loop
		for _, userEmail := range userEmails {
			checkEmailExist, err := i.repo.CheckUserByEmail(ctx, userEmail)
			if err != nil {
				return ServerError
			}
			if !checkEmailExist {
				return UserNotFound
			}
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
		if err = i.repo.AddFriendship(ctx, userEmails); err != nil {
			return ServerError
		}
		return nil
	})
}
