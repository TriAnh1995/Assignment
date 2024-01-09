package repository

import (
	"assignment/internal/repository/orm"
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (i RepoImplement) BlockNonSubscribedUser(ctx context.Context, userEmails []string) error {
	blockUser := &orm.Subscription{
		Requester: userEmails[0],
		Target:    userEmails[1],
		Status:    "blocked",
	}
	err := blockUser.Insert(ctx, i.pgConn, boil.Infer())
	return err
}
