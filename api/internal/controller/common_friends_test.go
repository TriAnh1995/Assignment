package controller

import (
	"assignment/internal/model"
	"assignment/internal/repository"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestControllerImplement_CommonFriends(t *testing.T) {
	type expectedCheckUserByEmail struct {
		expectedExist bool
		expectedErr   error
	}
	type expectedGetFriendsList struct {
		expectedList []string
		expectedErr  error
	}
	testCases := []struct {
		Name      string
		InputData []string
		expectedCheckUserByEmail
		expectedGetFriendsList1 expectedGetFriendsList
		expectedGetFriendsList2 expectedGetFriendsList
		expectedResult          model.FriendshipInfo
		expectedErr             error
	}{
		{
			Name:                     "Success",
			InputData:                []string{"user1@example.com", "user2@example.com"},
			expectedCheckUserByEmail: expectedCheckUserByEmail{true, nil},
			expectedGetFriendsList1:  expectedGetFriendsList{[]string{"friend1", "friend2", "friend4"}, nil},
			expectedGetFriendsList2:  expectedGetFriendsList{[]string{"friend1", "friend2", "friend3"}, nil},
			expectedResult: model.FriendshipInfo{
				Amounts: 2,
				List:    []string{"friend1", "friend2"},
			},
			expectedErr: nil,
		},
		{
			Name:                     "User not found",
			InputData:                []string{"user1@example.com", "user2@example.com"},
			expectedCheckUserByEmail: expectedCheckUserByEmail{false, nil},
			expectedGetFriendsList1:  expectedGetFriendsList{nil, nil},
			expectedGetFriendsList2:  expectedGetFriendsList{nil, nil},
			expectedResult:           model.FriendshipInfo{},
			expectedErr:              UserNotFound,
		},
		{
			Name:                     "Server error from CheckUserByEmail",
			InputData:                []string{"user1@example.com", "user2@example.com"},
			expectedCheckUserByEmail: expectedCheckUserByEmail{true, ServerError},
			expectedGetFriendsList1:  expectedGetFriendsList{nil, nil},
			expectedGetFriendsList2:  expectedGetFriendsList{nil, nil},
			expectedResult:           model.FriendshipInfo{},
			expectedErr:              ServerError,
		},
		{
			Name:                     "Server error from GetFriendsList",
			InputData:                []string{"user1@example.com", "user2@example.com"},
			expectedCheckUserByEmail: expectedCheckUserByEmail{true, nil},
			expectedGetFriendsList1:  expectedGetFriendsList{nil, ServerError},
			expectedGetFriendsList2:  expectedGetFriendsList{nil, nil},
			expectedResult:           model.FriendshipInfo{},
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
			for _, inputCase := range tc.InputData {
				repo.On("CheckUserByEmail", ctx, inputCase).
					Return(tc.expectedCheckUserByEmail.expectedExist, tc.expectedCheckUserByEmail.expectedErr)
			}

			repo.On("GetFriendsList", ctx, tc.InputData[0]).
				Return(tc.expectedGetFriendsList1.expectedList, tc.expectedGetFriendsList1.expectedErr)
			repo.On("GetFriendsList", ctx, tc.InputData[1]).
				Return(tc.expectedGetFriendsList2.expectedList, tc.expectedGetFriendsList2.expectedErr)
			// Run the test
			result, err := ctrl.CommonFriends(ctx, tc.InputData)
			// Check result
			if err != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, tc.expectedErr)
				require.Equal(t, tc.expectedResult, result)
			}
		})
	}
}
