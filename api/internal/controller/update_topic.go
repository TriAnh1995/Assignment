package controller

import (
	"assignment/internal/middleware"
	"assignment/internal/model"
	"context"
)

func (i CTRLImplement) UpdateTopic(ctx context.Context, updateInfo model.UpdateInfo) ([]string, error) {

	if err := i.checkEmail(ctx, updateInfo.Sender); err != nil {
		return []string{}, err
	}

	if err := i.repo.UpdateTopic(ctx, updateInfo); err != nil {
		return []string{}, middleware.ServerError
	}
	receivedUpdateList, err := i.repo.GetReceivedUpdateList(ctx, updateInfo.Sender)
	if err != nil {
		return []string{}, middleware.ServerError
	}
	finalList := append(receivedUpdateList, updateInfo.MentionedEmails...)

	return finalList, err

	//blockedList, err := i.repo.GetBlockedList(ctx, updateInfo.Sender)
	//if err != nil {
	//	return []string{}, ServerError
	//}
	//friendList, err := i.repo.GetFriendsList(ctx, updateInfo.Sender)
	//if err != nil {
	//	return []string{}, ServerError
	//}

	//
	//// Among three lists that got extract, we need the list of friends emails that didn't get blocked
	//blockedFriendList := ExtractCommonElements(blockedList, friendList)
	//nonBlockedFriendList := RemoveElements(friendList, blockedFriendList)
	//
	//// And a unique emails in follower and friends circle
	//uniqueFollowerAndFriends := CombineUnique(followedList, nonBlockedFriendList)
	//
	//// Combine with the emails mentioned in the text, we got the final list
	//finalList := append(uniqueFollowerAndFriends, updateInfo.MentionedEmails...)

}
