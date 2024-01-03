package repository

import (
	"assignment/internal/repository/orm"
	"context"
)

func (i RepoImplement) CheckIfFollowed(ctx context.Context, userEmails []string) (bool, error) {
	exist, err := orm.Subscriptions(
		orm.SubscriptionWhere.Requester.EQ(userEmails[0]),
		orm.SubscriptionWhere.Target.EQ(userEmails[1]),
		orm.SubscriptionWhere.Status.EQ("followed"),
	).Exists(ctx, i.pgConn)
	return exist, err
}
