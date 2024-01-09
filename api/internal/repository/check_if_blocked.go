package repository

import (
	"assignment/internal/repository/orm"
	"context"
)

func (i RepoImplement) CheckIfBlocked(ctx context.Context, userEmails []string) (bool, error) {

	exists, err := orm.Subscriptions(
		orm.SubscriptionWhere.Requester.EQ(userEmails[0]),
		orm.SubscriptionWhere.Target.EQ(userEmails[1]),
		orm.SubscriptionWhere.Status.EQ("blocked"),
	).Exists(ctx, i.pgConn)
	return exists, err
}
