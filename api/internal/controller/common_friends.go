package controller

import (
	"assignment/internal/middleware"
	"assignment/internal/model"
	"context"
)

func (i CTRLImplement) CommonFriends(ctx context.Context, userEmails []string) (common model.FriendshipInfo, err error) {
	// CheckEmail
	if err = i.checkEmail(ctx, userEmails[0]); err != nil {
		return model.FriendshipInfo{}, err
	}
	if err = i.checkEmail(ctx, userEmails[1]); err != nil {
		return model.FriendshipInfo{}, err
	}
	// Extract list of friends for both users
	listOfFriends1, err := i.repo.GetFriendsList(ctx, userEmails[0])
	if err != nil {
		return model.FriendshipInfo{}, middleware.ServerError
	}
	listOfFriends2, err := i.repo.GetFriendsList(ctx, userEmails[1])
	if err != nil {
		return model.FriendshipInfo{}, middleware.ServerError
	}

	friendsInCommon := make(map[string]bool)
	for _, friend := range listOfFriends1 {
		friendsInCommon[friend] = true
	}
	for _, friend := range listOfFriends2 {
		if friendsInCommon[friend] {
			common.List = append(common.List, friend)
		}
	}
	common.Amounts = len(common.List)

	return common, nil
}
