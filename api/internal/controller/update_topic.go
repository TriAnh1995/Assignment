package controller

import (
	"assignment/internal/model"
	"context"
)

func (i CTRLImplement) UpdateTopic(ctx context.Context, updateInfo model.UpdateInfo) ([]string, error) {
	// Check email exist
	checkEmailExist, err := i.repo.CheckUserByEmail(ctx, updateInfo.Sender)
	if err != nil {
		return []string{}, ServerError
	}
	if !checkEmailExist {
		return []string{}, UserNotFound
	}

	// Insert the Topic status and Update Message
	if err = i.repo.UpdateTopic(ctx, updateInfo); err != nil {
		return []string{}, ServerError
	}

	// Extract list of blocked users and list of friends that the Sender have
	blockedList, err := i.repo.GetBlockedList(ctx, updateInfo.Sender)
	if err != nil {
		return []string{}, ServerError
	}

	friendList, err := i.repo.GetFriendsList(ctx, updateInfo.Sender)
	if err != nil {
		return []string{}, ServerError
	}

	// Using map string to find the list of friends that didn't get blocked
	var nonBlockedFriends []string
	Map := make(map[string]bool)
	for _, username := range friendList {
		Map[username] = true
	}
	for _, friend := range blockedList {
		if !Map[friend] {
			nonBlockedFriends = append(nonBlockedFriends, friend)
		}
	}

	// Extract the list of follower from subscription db
	followedList, err := i.repo.GetFollowerList(ctx, updateInfo.Sender)
	if err != nil {
		return []string{}, ServerError
	}

	// Create a unique map to store elements from the nonBlockedFriends and followedList
	uniqueMap := make(map[string]struct{})
	for _, s := range nonBlockedFriends {
		uniqueMap[s] = struct{}{}
	}
	for _, s := range followedList {
		uniqueMap[s] = struct{}{}
	}
	// Create a new slice to hold the unique emails from nonBlockedFriends and followedList
	updateReceivedList := make([]string, 0, len(uniqueMap))
	for s := range uniqueMap {
		updateReceivedList = append(updateReceivedList, s)
	}
	return updateReceivedList, nil
}
