package repository

import (
	"assignment/internal/repository/orm"
	"context"
)

// GetFollowedList retrieve a list of ...
func (i RepoImplement) GetFollowedList(ctx context.Context, userEmail string) ([]string, error) {
	followed, err := orm.Subscriptions(
		orm.SubscriptionWhere.Target.EQ(userEmail),
		orm.SubscriptionWhere.Status.EQ("followed"),
	).All(ctx, i.pgConn)
	if err != nil {
		return []string{}, err
	}

	listOfFollowed := make([]string, len(followed))
	for k, follow := range followed {
		listOfFollowed[k] = follow.Requester
	}
	return listOfFollowed, err
}
