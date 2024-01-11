package controller

import (
	"assignment/internal/middleware"
	"assignment/internal/model"
	"assignment/internal/repository"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestControllerImplement_FriendsList(t *testing.T) {
	type expectedCheckUserByEmail struct {
		expectedExist bool
		expectedErr   error
	}
	type expectedGetFriendsList struct {
		expectedList []string
		expectedErr  error
	}
	testCases := []struct {
		Name       string
		InputEmail string
		expectedCheckUserByEmail
		expectedGetFriendsList
		expectedResult model.FriendshipInfo
		expectedErr    error
	}{
		{
			Name:                     "Success",
			InputEmail:               "user@example.com",
			expectedCheckUserByEmail: expectedCheckUserByEmail{true, nil},
			expectedGetFriendsList:   expectedGetFriendsList{[]string{"friend1", "friend2", "friend3", "friend4", "friend5"}, nil},
			expectedResult: model.FriendshipInfo{
				Amounts: 5,
				List:    []string{"friend1", "friend2", "friend3", "friend4", "friend5"},
			},
			expectedErr: nil,
		},
		{
			Name:                     "User Email not found",
			InputEmail:               "user@example.com",
			expectedCheckUserByEmail: expectedCheckUserByEmail{false, nil},
			expectedGetFriendsList:   expectedGetFriendsList{nil, nil},
			expectedResult:           model.FriendshipInfo{},
			expectedErr:              middleware.UserNotFound,
		},
		{
			Name:                     "Server error from CheckUserByEmail",
			InputEmail:               "user@example.com",
			expectedCheckUserByEmail: expectedCheckUserByEmail{true, middleware.ServerError},
			expectedGetFriendsList:   expectedGetFriendsList{nil, nil},
			expectedResult:           model.FriendshipInfo{},
			expectedErr:              middleware.ServerError,
		},
		{
			Name:                     "Server error from GetFriendsList",
			InputEmail:               "user@example.com",
			expectedCheckUserByEmail: expectedCheckUserByEmail{true, nil},
			expectedGetFriendsList:   expectedGetFriendsList{nil, middleware.ServerError},
			expectedResult:           model.FriendshipInfo{},
			expectedErr:              middleware.ServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			// Setup Instance
			repo := &repository.MockRepository{}
			ctrl := New(repo)
			ctx := context.Background()

			repo.On("CheckUserByEmail", ctx, tc.InputEmail).
				Return(tc.expectedCheckUserByEmail.expectedExist, tc.expectedCheckUserByEmail.expectedErr)

			repo.On("GetFriendsList", ctx, tc.InputEmail).
				Return(tc.expectedGetFriendsList.expectedList, tc.expectedGetFriendsList.expectedErr)

			result, err := ctrl.FriendsList(ctx, tc.InputEmail)

			if err != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, tc.expectedErr)
				require.Equal(t, tc.expectedResult, result)
			}
		})
	}
}
