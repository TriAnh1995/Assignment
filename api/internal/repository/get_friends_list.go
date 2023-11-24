package repository

import (
	"assignment/internal/repository/orm"
	"context"
)

// GetFriendsList GetFriendList retrieve a list of friend names a user have
func (i RepoImplement) GetFriendsList(ctx context.Context, userEmail string) ([]string, error) {
	var listOfFriends []string
	friends1, err := orm.Friendships(
		orm.FriendshipWhere.UserEmail1.EQ(userEmail)).All(ctx, i.pgConn)
	if err != nil {
		return []string{}, err
	}
	for _, friend := range friends1 {
		listOfFriends = append(listOfFriends, friend.UserEmail2)
	}

	friends2, err := orm.Friendships(
		orm.FriendshipWhere.UserEmail2.EQ(userEmail)).All(ctx, i.pgConn)
	if err != nil {
		return []string{}, err
	}
	for _, friend := range friends2 {
		listOfFriends = append(listOfFriends, friend.UserEmail1)
	}
	return listOfFriends, err
}
