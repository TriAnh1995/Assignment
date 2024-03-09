package repository

import (
	"assignment/internal/repository/orm"
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (i RepoImplement) SubscribeToNonBlocked(ctx context.Context, emails []string) error {
	subscribe := &orm.Subscription{
		Requester: emails[0],
		Target:    emails[1],
		Status:    orm.StatusTypeFollowed,
	}

	err := subscribe.Insert(ctx, i.pgConn, boil.Infer())
	return err
}
