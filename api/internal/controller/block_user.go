package controller

import (
	"assignment/internal/middleware"
	"context"

	"gorm.io/gorm"
)

func (i CTRLImplement) BlockUsers(ctx context.Context, userEmails []string) error {
	tx, err := middleware.GormConnection()
	if err != nil {
		return middleware.ServerError
	}
	return tx.Transaction(func(*gorm.DB) error {
		if err = i.checkEmail(ctx, userEmails[0]); err != nil {
			return err
		}
		if err = i.checkEmail(ctx, userEmails[1]); err != nil {
			return err
		}

		checkSubscription, err := i.repo.CheckSubscription(ctx, userEmails)
		if err != nil {
			return middleware.ServerError
		}
		switch checkSubscription {
		case "Blocked":
			{
				return middleware.AlreadyBlocked
			}
		case "Followed", "Default":
			{
				err = i.repo.BlockUserInsideSubscription(ctx, userEmails)
				if err != nil {
					return middleware.ServerError
				}
			}
		case "None":
			{
				err = i.repo.BlockNonSubscribedUser(ctx, userEmails)
				if err != nil {
					return middleware.ServerError
				}
			}
		}
		return nil
	})
}
