package repository

import (
	"assignment/internal/repository/orm"
	"context"
)

// GetFriendsList retrieve a list of friend names a user have
func (i RepoImplement) GetFriendsList(ctx context.Context, userEmail string) ([]string, error) {
	//Create a slice to store the list of friends for the userEmail.
	var listOfFriends []string
	//Use the 'All' method to extract all the rows that contain the userEmail in the first column.
	friends1, err := orm.Friendships(
		orm.FriendshipWhere.UserEmail1.EQ(userEmail)).All(ctx, i.pgConn)
	if err != nil {
		return []string{}, err
	}
	//Store the emails of userEmail's friends, which are located in the second column, in the list.
	for _, friend := range friends1 {
		listOfFriends = append(listOfFriends, friend.UserEmail2)
	}
	//Repeat the process for the other case where the userEmail is contained in the second column.
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
