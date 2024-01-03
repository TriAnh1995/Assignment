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
	// Put both list in function to find the common friends among two list of friends
	common.List = ExtractCommonElements(listOfFriends1, listOfFriends2)
	common.Amounts = len(common.List)
	return common, nil
}
