package controller

import (
	"context"

	"gorm.io/gorm"
)

func (i CTRLImplement) AddFriends(ctx context.Context, tx *gorm.DB, userEmails []string) error {
	// Start Business Logic
	// Check whether ot not both users exists in db with for loop
	for _, userEmail := range userEmails {
		checkEmailExist, err := i.repo.CheckUserByEmail(ctx, userEmail)
		if err != nil {
			tx.Rollback()
			return ServerError
		}
		if !checkEmailExist {
			tx.Rollback()
			return UserNotFound
		}
	}

	// Check if both user already friends in db
	checkFriendshipExist, err := i.repo.CheckFriendship(ctx, userEmails)
	if err != nil {
		tx.Rollback()
		return ServerError
	}
	if checkFriendshipExist {
		tx.Rollback()
		return FriendshipExisted
	}

	// Add friendship to db
	if err = i.repo.AddFriendship(ctx, userEmails); err != nil {
		tx.Rollback()
		return ServerError
	}

	// Commit to transaction
	tx.Commit()
	return nil
}
