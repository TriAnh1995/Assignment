package repository

import (
	"assignment/internal/repository/orm"
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// AddFriendship insert two usernames to friendship table in db
func (i RepoImplement) AddFriendship(ctx context.Context, firstEmail, secondEmail string) error {
	createFriendShip := &orm.Friendship{
		UserEmail1: firstEmail,
		UserEmail2: secondEmail,
	}

	err := createFriendShip.Insert(ctx, i.pgConn, boil.Infer())
	return err
}
