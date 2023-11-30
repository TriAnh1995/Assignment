package repository

import (
	"assignment/internal/repository/orm"
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// AddFriendship insert two usernames to friendship table in db
func (i RepoImplement) AddFriendship(ctx context.Context, userEmails []string) error {
	createFriendShip := &orm.Friendship{
		UserEmail1: userEmails[0],
		UserEmail2: userEmails[1],
	}
	return createFriendShip.Insert(ctx, i.pgConn, boil.Infer())
}
