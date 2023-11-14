package controller

import (
	"context"

	_ "github.com/lib/pq"
)

func (i CTRLImplement) AddFriend(ctx context.Context, userEmails []string) error {
	tx, err := ConnectToDatabase()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
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

	tx.Commit()
	return nil

}
