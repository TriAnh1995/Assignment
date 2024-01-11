package controller

import (
	"assignment/internal/middleware"
	"assignment/internal/model"
	"context"
)

func (i CTRLImplement) FriendsList(ctx context.Context, userEmail string) (model.FriendshipInfo, error) {
	// Check whether user already in database through email
	err := i.checkEmail(ctx, userEmail)
	if err != nil {
		return model.FriendshipInfo{}, err
	}
	//Using FriendshipInfo model for the output
	var FriendsInfo model.FriendshipInfo
	//Get the list of friends from repository
	if FriendsInfo.List, err = i.repo.GetFriendsList(ctx, userEmail); err != nil {
		return model.FriendshipInfo{}, middleware.ServerError
	}
	//The length of the list is the number of friends
	FriendsInfo.Amounts = len(FriendsInfo.List)

	return FriendsInfo, nil
}
