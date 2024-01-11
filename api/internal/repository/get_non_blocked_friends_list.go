package repository

import (
	"assignment/internal/repository/orm"
	"context"
)

// GetFriendsList retrieve a list of friend names a user have
func (i RepoImplement) GetNonBlocList(ctx context.Context, userEmail string) ([]string, error) {
	var listOfFriends []string

	friends, err := orm.Relationships(
		orm.RelationshipWhere.UserEmail1.EQ(userEmail),
		orm.RelationshipWhere.Friendship.EQ("friend"),
	).All(ctx, i.pgConn)
	if err != nil {
		return []string{}, err
	}

	for _, friend := range friends {
		listOfFriends = append(listOfFriends, friend.UserEmail2)
	}

	return listOfFriends, err
}
