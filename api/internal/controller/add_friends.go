package controller

import (
	"context"

	"gorm.io/gorm"
)

func (i CTRLImplement) AddFriend(ctx context.Context, userEmails []string) error {
	// Begin Transaction
	err := i.gormDB.Transaction(func(*gorm.DB) error {
		// Business Logic...
		for _, userEmail := range userEmails {
			checkEmailExist, err := i.repo.CheckUserByEmail(ctx, userEmail)
			if err != nil {
				return ServerError
			}
			if !checkEmailExist {
				return UserNotFound
			}
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
		// ... Commit Transaction
		if err = i.gormDB.Commit().Error; err != nil {
			return ServerError
		}
		return nil
	})
	return err
}
