package controller

import (
	"context"

	"gorm.io/gorm"
)

func (i CTRLImplement) AddFriend(ctx context.Context, userEmails []string) error {
	newTransaction, err := GormConnection()
	if err != nil {
		return ServerError
	}
	return newTransaction.Transaction(func(*gorm.DB) error {
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
		if err = newTransaction.Commit().Error; err != nil {
			return ServerError
		}
		return nil
	})
}
