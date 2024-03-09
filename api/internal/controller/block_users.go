package controller

import (
	"context"

	"gorm.io/gorm"
)

func (i CTRLImplement) BlockUsers(ctx context.Context, userEmails []string) error {

	tx, err := ConnectToDatabase()
	if err != nil {
		return ServerError
	}
	return tx.Transaction(func(*gorm.DB) error {
		if err = i.checkEmail(ctx, userEmails[0]); err != nil {
			return err
		}
		if err = i.checkEmail(ctx, userEmails[1]); err != nil {
			return err
		}

		checkIfBlocked, err := i.repo.CheckIfBlocked(ctx, userEmails)
		if err != nil {
			return ServerError
		}
		if checkIfBlocked {
			return AlreadyBlocked
		}

		checkIfFollowed, err := i.repo.CheckIfFollowed(ctx, userEmails)
		if err != nil {
			return ServerError
		}

		if checkIfFollowed {
			err = i.repo.BlockToSubscribed(ctx, userEmails)
			if err != nil {
				return ServerError
			}
		} else {
			if err = i.repo.BlockToNonSubscribed(ctx, userEmails); err != nil {
				return ServerError
			}
		}
		return nil
	})
}
