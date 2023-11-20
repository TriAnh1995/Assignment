package controller

import (
	"assignment/internal/model"
	"context"
)

func (i CTRLImplement) FriendsList(ctx context.Context, userEmail string) (model.FriendshipInfo, error) {
	// Check whether user already in database through email
	userExists, err := i.repo.CheckUserByEmail(ctx, userEmail)
	if err != nil {
		return model.FriendshipInfo{}, ServerError
	}
	if !userExists {
		return model.FriendshipInfo{}, UserNotFound
	}

	var FriendsInfo model.FriendshipInfo

	if FriendsInfo.List, err = i.repo.GetFriendsList(ctx, userEmail); err != nil {
		return model.FriendshipInfo{}, ServerError
	}
	FriendsInfo.Amounts = len(FriendsInfo.List)

	return FriendsInfo, nil
}
