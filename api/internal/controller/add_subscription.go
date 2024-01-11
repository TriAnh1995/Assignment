package controller

import (
	"assignment/internal/middleware"
	"context"

	"gorm.io/gorm"
)

func (i CTRLImplement) Subscription(ctx context.Context, userEmails []string) error {
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
		case "Followed":
			{
				return middleware.AlreadySubscribed
			}
		case "Blocked", "Default":
			{
				err = i.repo.SubscribeToBlocked(ctx, userEmails)
				if err != nil {
					return middleware.ServerError
				}
			}
		case "None":
			{
				err = i.repo.AddSubscriptions(ctx, userEmails)
				if err != nil {
					return middleware.ServerError
				}
			}
		}
		return nil
	})
}
