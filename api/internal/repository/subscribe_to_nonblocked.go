package repository

import (
	"assignment/internal/middleware"
	"assignment/internal/repository/orm"
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"gorm.io/gorm"
)

func (i RepoImplement) SubscribeToNonBlocked(ctx context.Context, emails []string) error {
	tx, err := middleware.GormConnection()
	if err != nil {
		return err
	}
	return tx.Transaction(func(db *gorm.DB) error {
		createFriendShip := &orm.Relationship{
			UserEmail1:   emails[0],
			UserEmail2:   emails[1],
			Subscription: "followed",
			Friendship:   "stranger",
		}
		err = createFriendShip.Insert(ctx, i.pgConn, boil.Infer())

		if err != nil {
			return err
		}
		createFriendShip = &orm.Relationship{
			UserEmail1:   emails[1],
			UserEmail2:   emails[0],
			Subscription: "default",
			Friendship:   "stranger",
		}
		err = createFriendShip.Insert(ctx, i.pgConn, boil.Infer())
		if err != nil {
			return err
		}
		return nil
	})
}
