package repository

import (
	"assignment/internal/repository/orm"
	"context"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetFriendsList retrieve a list of friend names a user have
func (i RepoImplement) GetFriendsList(ctx context.Context, userEmail string) ([]string, error) {
	// Using Query to find userEmail on any columns in database
	friendships, err := orm.Friendships(
		qm.Where(orm.FriendshipColumns.UserEmail1+" = ? OR "+orm.FriendshipColumns.UserEmail2+" = ?", userEmail, userEmail)).
		All(ctx, i.pgConn)
	if err != nil {
		return []string{}, err
	}
	// Put the Email on the column that have different value of userEmail to the list
	var friendsList []string
	for _, friendship := range friendships {
		if friendship.UserEmail1 == userEmail {
			friendsList = append(friendsList, friendship.UserEmail2)
		} else {
			friendsList = append(friendsList, friendship.UserEmail1)
		}
	}
	return friendsList, nil
}
