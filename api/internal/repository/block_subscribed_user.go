package repository

import (
	"assignment/internal/repository/orm"
	"context"
)

func (i RepoImplement) BlockUserInsideSubscription(ctx context.Context, userEmails []string) error {
	_, err := orm.Subscriptions(
		orm.SubscriptionWhere.Requester.EQ(userEmails[0]),
		orm.SubscriptionWhere.Target.EQ(userEmails[1]),
	).UpdateAll(ctx, i.pgConn, orm.M{"status": "blocked"})

	return err
}
