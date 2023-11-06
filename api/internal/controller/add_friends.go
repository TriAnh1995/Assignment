package controller

import (
	"context"
)

func (i CTRLImplement) AddFriends(ctx context.Context, userEmails []string) error {
	// Establish database connection
	appDB, err := ConnectToDatabase()
	if err != nil {
		return ServerError
	}
	// Start a transaction
	tx, err := appDB.BeginTx(ctx, nil)
	if err != nil {
		return ServerError
	}
	defer tx.Rollback()
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
	// Commit to transaction
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
