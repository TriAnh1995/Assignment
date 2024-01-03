package controller

import (
	"assignment/internal/model"
	"context"
)

func (i CTRLImplement) FriendsList(ctx context.Context, userEmail string) (model.FriendshipInfo, error) {
	//Using FriendshipInfo model for the output
	var FriendsInfo model.FriendshipInfo
	var err error
	// Check whether user already in database through email
	if err = i.checkEmail(ctx, userEmail); err != nil {
		return model.FriendshipInfo{}, err
	}
	//Get the list of friends from repository
	if FriendsInfo.List, err = i.repo.GetFriendsList(ctx, userEmail); err != nil {
		return model.FriendshipInfo{}, ServerError
	}
	//The length of the list is the number of friends
	FriendsInfo.Amounts = len(FriendsInfo.List)

	return FriendsInfo, nil
}
