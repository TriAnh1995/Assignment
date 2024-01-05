package controller

import (
	"assignment/internal/model"
	"context"
)

func (i CTRLImplement) FriendsList(ctx context.Context, userEmail string) (friendsInfo model.FriendshipInfo, err error) {
	// Check whether user already in database through email
	if err = i.checkEmail(ctx, userEmail); err != nil {
		return model.FriendshipInfo{}, err
	}

	//Get the list of friends from repository
	if friendsInfo.List, err = i.repo.GetFriendsList(ctx, userEmail); err != nil {
		return model.FriendshipInfo{}, ServerError
	}

	//The length of the list is the number of friends
	friendsInfo.Amounts = len(friendsInfo.List)

	return friendsInfo, nil
}
