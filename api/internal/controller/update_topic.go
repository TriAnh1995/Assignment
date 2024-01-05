package controller

import (
	"assignment/internal/model"
	"context"
)

func (i CTRLImplement) UpdateTopic(ctx context.Context, updateInfo model.UpdateInfo) ([]string, error) {
	// Check email exist
	if err := i.checkEmail(ctx, updateInfo.Sender); err != nil {
		return []string{}, err
	}

	// Insert the Topic status and Update Message
	if err := i.repo.UpdateTopic(ctx, updateInfo); err != nil {
		return []string{}, ServerError
	}

	// Extract list of blocked users, followed user and friends of the Sender
	blockedList, err := i.repo.GetBlockedList(ctx, updateInfo.Sender)
	if err != nil {
		return []string{}, ServerError
	}
	friendList, err := i.repo.GetFriendsList(ctx, updateInfo.Sender)
	if err != nil {
		return []string{}, ServerError
	}
	followedList, err := i.repo.GetFollowerList(ctx, updateInfo.Sender)
	if err != nil {
		return []string{}, ServerError
	}

	// Among three lists that got extract, we need the list of friends emails that didn't get blocked
	blockedFriendList := ExtractCommonElements(blockedList, friendList)
	nonBlockedFriendList := RemoveElements(friendList, blockedFriendList)

	// And a unique emails in follower and friends circle
	commonInFollowersAndFriends := ExtractCommonElements(followedList, friendList)
	combinedFriendsAndFollower := append(nonBlockedFriendList, followedList...)
	uniqueFollowerAndFriends := RemoveElements(combinedFriendsAndFollower, commonInFollowersAndFriends)

	// Combine with the emails mentioned in the text, we got the final list
	finalList := append(uniqueFollowerAndFriends, updateInfo.MentionedEmail)
	return finalList, nil
}
