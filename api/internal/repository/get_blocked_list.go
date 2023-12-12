package repository

import (
	"assignment/internal/repository/orm"
	"context"
)

// GetBlockedList retrieve a list of ...
func (i RepoImplement) GetBlockedList(ctx context.Context, userEmail string) ([]string, error) {
	blocked, err := orm.Subscriptions(
		orm.SubscriptionWhere.Requester.EQ(userEmail),
		orm.SubscriptionWhere.Status.EQ("blocked"),
	).All(ctx, i.pgConn)
	if err != nil {
		return []string{}, err
	}

	listOfBlocked := make([]string, len(blocked))
	for k, block := range blocked {
		listOfBlocked[k] = block.Target
	}
	return listOfBlocked, err
}
