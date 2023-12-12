package controller

import (
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
		return model.FriendshipInfo{}, ServerError
	}
	listOfFriends2, err := i.repo.GetFriendsList(ctx, userEmails[1])
	if err != nil {
		return model.FriendshipInfo{}, ServerError
	}

	// Create an empty list of common friends
	friendsInCommon := make(map[string]bool)

	// If the friend exist in the list of one person...
	for _, friend := range listOfFriends1 {
		friendsInCommon[friend] = true
	}
	// ... and the other...
	for _, friend := range listOfFriends2 {
		if friendsInCommon[friend] {
			// ... That's the common friend. So we insert them into the list
			common.List = append(common.List, friend)
		}
	}
	// The number of common friend is the length of the list
	common.Amounts = len(common.List)

	return common, nil
}
