package repository

import (
	"assignment/internal/middleware"
	"assignment/internal/repository/orm"
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"gorm.io/gorm"
)

// AddFriendship insert two usernames to friendship table in db
func (i RepoImplement) AddFriendship(ctx context.Context, userEmails []string) error {
	tx, err := middleware.GormConnection()
	if err != nil {
		return err
	}
	return tx.Transaction(func(db *gorm.DB) error {
		for _, user := range []string{userEmails[0], userEmails[1]} {
			createFriendShip := &orm.Relationship{
				UserEmail1:   userEmails[0],
				UserEmail2:   userEmails[1],
				Subscription: "default",
				Friendship:   "friend",
			}

			if user != userEmails[0] {
				createFriendShip.UserEmail1, createFriendShip.UserEmail2 = createFriendShip.UserEmail2, createFriendShip.UserEmail1
			}

			if err := createFriendShip.Insert(ctx, i.pgConn, boil.Infer()); err != nil {
				return err
			}
		}
		return nil
	})
}
