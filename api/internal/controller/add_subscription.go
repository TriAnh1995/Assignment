package controller

import (
	"context"

	"gorm.io/gorm"
)

func (i CTRLImplement) AddSubscription(ctx context.Context, userEmails []string) error {
	tx, err := ConnectToDatabase()
	if err != nil {
		return ServerError
	}
	return tx.Transaction(func(*gorm.DB) error {
		// Check Emails
		if err = i.checkEmail(ctx, userEmails[0]); err != nil {
			return err
		}
		if err = i.checkEmail(ctx, userEmails[1]); err != nil {
			return err
		}

		// Check Relationship between two emails
		checkIfFollowed, err := i.repo.CheckIfFollowed(ctx, userEmails)
		if err != nil {
			return ServerError
		}
		if checkIfFollowed {
			return AlreadySubscribed
		}

		checkIfBlocked, err := i.repo.CheckIfBlocked(ctx, userEmails)
		if err != nil {
			return ServerError
		}

		if checkIfBlocked {
			err = i.repo.SubscribeToBlocked(ctx, userEmails)
			if err != nil {
				return ServerError
			}
		} else {
			if err = i.repo.SubscribeToNonBlocked(ctx, userEmails); err != nil {
				return ServerError
			}
		}
		return nil
	})
}
