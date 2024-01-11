package repository

import (
	"assignment/internal/repository/orm"
	"context"
)

func (i RepoImplement) GetReceivedUpdateList(ctx context.Context, userEmail string) ([]string, error) {
	var ReceivedUpdateList []string

	followers, err := orm.Relationships(
		orm.RelationshipWhere.UserEmail2.EQ(userEmail),
		orm.RelationshipWhere.Subscription.EQ("followed"),
	).All(ctx, i.pgConn)
	if err != nil {
		return []string{}, err
	}

	friends, err := orm.Relationships(
		orm.RelationshipWhere.UserEmail2.EQ(userEmail),
		orm.RelationshipWhere.Friendship.EQ("friend"),
		orm.RelationshipWhere.Subscription.EQ("default"),
	).All(ctx, i.pgConn)
	if err != nil {
		return []string{}, err
	}

	for _, follower := range followers {
		ReceivedUpdateList = append(ReceivedUpdateList, follower.UserEmail1)
	}
	for _, friend := range friends {
		ReceivedUpdateList = append(ReceivedUpdateList, friend.UserEmail1)
	}

	return ReceivedUpdateList, err
}
