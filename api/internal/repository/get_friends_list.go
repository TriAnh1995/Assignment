package repository

import (
	"assignment/internal/repository/orm"
	"context"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetFriendsList GetFriendList retrieve a list of friend names a user have
func (i RepoImplement) GetFriendsList(ctx context.Context, userEmail string) ([]string, error) {
	// Using method `All` from SQLBoiler to extract all the rows in 'friendship' table that contain userEmail
	friends, err := orm.Friendships(
		qm.Where("user_email_1 = ? OR user_email_2 = ?", userEmail, userEmail)).
		All(ctx, i.pgConn)
	if err != nil {
		return []string{}, err
	}

	// Create a slice to store list of friends
	listOfFriends := make([]string, len(friends))

	// Put all the rows in 'for' loop, check for each column which have emails that not userEmail and store them in the list
	for n, friend := range friends {
		if friend.UserEmail1 != userEmail {
			listOfFriends[n] = friend.UserEmail1
		} else {
			listOfFriends[n] = friend.UserEmail2
		}
	}
	return listOfFriends, err
}
