package controller

import (
	"assignment/internal/model"
	"assignment/internal/repository"
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCTRLImplement_UpdateTopic(t *testing.T) {
	type expectedCheckUserByEmail struct {
		expectedExist bool
		expectedErr   error
	}
	type expectedGetFriendsList struct {
		expectedList []string
		expectedErr  error
	}
	type expectedGetBlockedList struct {
		expectedList []string
		expectedErr  error
	}
	type expectedGetFollowerList struct {
		expectedList []string
		expectedErr  error
	}
	testCases := []struct {
		Name      string
		InputData model.UpdateInfo
		expectedCheckUserByEmail
		expectedUpdateTopic error
		expectedGetFriendsList
		expectedGetBlockedList
		expectedGetFollowerList
		expectedResult []string
		expectedErr    error
	}{
		{
			Name: "Success",
			InputData: model.UpdateInfo{
				Sender: "bob@example.com",
				Text:   "Hello World! C",
			},
			expectedCheckUserByEmail: expectedCheckUserByEmail{true, nil},
			expectedUpdateTopic:      nil,
			expectedGetFriendsList:   expectedGetFriendsList{[]string{"A", "B", "D"}, nil},
			expectedGetBlockedList:   expectedGetBlockedList{[]string{"A", "F", "E"}, nil},
			expectedGetFollowerList:  expectedGetFollowerList{[]string{"D", "J", "G"}, nil},
			expectedResult:           []string{"J", "G", "F", "E", "D", "B"},
			expectedErr:              nil,
		},
		{
			Name: "Server error from CheckUserByEmail",
			InputData: model.UpdateInfo{
				Sender: "bob@example.com",
				Text:   "Hello World! C",
			},
			expectedCheckUserByEmail: expectedCheckUserByEmail{true, ServerError},
			expectedUpdateTopic:      nil,
			expectedGetFriendsList:   expectedGetFriendsList{},
			expectedGetBlockedList:   expectedGetBlockedList{},
			expectedGetFollowerList:  expectedGetFollowerList{},
			expectedResult:           []string{},
			expectedErr:              ServerError,
		},
		{
			Name: "User Email not found",
			InputData: model.UpdateInfo{
				Sender: "bob@example.com",
				Text:   "Hello World! C",
			},
			expectedCheckUserByEmail: expectedCheckUserByEmail{false, nil},
			expectedUpdateTopic:      nil,
			expectedGetFriendsList:   expectedGetFriendsList{},
			expectedGetBlockedList:   expectedGetBlockedList{},
			expectedGetFollowerList:  expectedGetFollowerList{},
			expectedResult:           []string{},
			expectedErr:              UserNotFound,
		},
		{
			Name: "Server error from UpdateTopic",
			InputData: model.UpdateInfo{
				Sender: "bob@example.com",
				Text:   "Hello World! C",
			},
			expectedCheckUserByEmail: expectedCheckUserByEmail{true, nil},
			expectedUpdateTopic:      ServerError,
			expectedGetFriendsList:   expectedGetFriendsList{},
			expectedGetBlockedList:   expectedGetBlockedList{},
			expectedGetFollowerList:  expectedGetFollowerList{},
			expectedResult:           []string{},
			expectedErr:              ServerError,
		},
		{
			Name:                     "Server error from GetFriendsList",
			InputData:                model.UpdateInfo{},
			expectedCheckUserByEmail: expectedCheckUserByEmail{true, nil},
			expectedUpdateTopic:      nil,
			expectedGetFriendsList:   expectedGetFriendsList{[]string{}, ServerError},
			expectedGetBlockedList:   expectedGetBlockedList{[]string{}, nil},
			expectedGetFollowerList:  expectedGetFollowerList{},
			expectedResult:           nil,
			expectedErr:              ServerError,
		},
		{
			Name:                     "Server error from GetBlockedList",
			InputData:                model.UpdateInfo{},
			expectedCheckUserByEmail: expectedCheckUserByEmail{true, nil},
			expectedUpdateTopic:      nil,
			expectedGetFriendsList:   expectedGetFriendsList{[]string{}, nil},
			expectedGetBlockedList:   expectedGetBlockedList{[]string{}, ServerError},
			expectedGetFollowerList:  expectedGetFollowerList{},
			expectedResult:           nil,
			expectedErr:              ServerError,
		},
		{
			Name:                     "Server error from GetFollowerList",
			InputData:                model.UpdateInfo{},
			expectedCheckUserByEmail: expectedCheckUserByEmail{true, nil},
			expectedUpdateTopic:      nil,
			expectedGetFriendsList:   expectedGetFriendsList{[]string{}, nil},
			expectedGetBlockedList:   expectedGetBlockedList{[]string{}, nil},
			expectedGetFollowerList:  expectedGetFollowerList{[]string{}, ServerError},
			expectedResult:           nil,
			expectedErr:              ServerError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Setup Instance
			repo := &repository.MockRepository{}
			ctrl := New(repo)
			ctx := context.Background()

			// Setup mock and mock behavior
			repo.On("CheckUserByEmail", ctx, tc.InputData.Sender).Return(tc.expectedCheckUserByEmail.expectedExist, tc.expectedCheckUserByEmail.expectedErr)
			repo.On("UpdateTopic", ctx, tc.InputData).Return(tc.expectedUpdateTopic)
			repo.On("GetBlockedList", ctx, tc.InputData.Sender).Return(tc.expectedGetBlockedList.expectedList, tc.expectedGetBlockedList.expectedErr)
			repo.On("GetFriendsList", ctx, tc.InputData.Sender).Return(tc.expectedGetFriendsList.expectedList, tc.expectedGetFriendsList.expectedErr)
			repo.On("GetFollowerList", ctx, tc.InputData.Sender).Return(tc.expectedGetFollowerList.expectedList, tc.expectedGetFollowerList.expectedErr)

			// Run the test
			result, err := ctrl.UpdateTopic(ctx, tc.InputData)
			// Check result
			if err != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, tc.expectedErr)
				reflect.DeepEqual(tc.expectedResult, result)
			}
		})
	}
}
