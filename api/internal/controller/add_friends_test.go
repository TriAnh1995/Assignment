package controller

import (
	"assignment/internal/repository"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestControllerImplement_AddFriend(t *testing.T) {
	type expectedCheckUserByEmail struct {
		expectedExist bool
		expectedErr   error
	}
	type expectedCheckFriendship struct {
		expectedExist bool
		expectedErr   error
	}
	testCases := []struct {
		Name  string
		Input []string
		expectedCheckUserByEmail
		expectedCheckFriendship
		expectedAddFriendship error
		expectedErr           error
	}{
		{
			Name:                     "Success",
			Input:                    []string{"firstuser@example.com", "seconduser@example.com"},
			expectedCheckUserByEmail: expectedCheckUserByEmail{true, nil},
			expectedCheckFriendship:  expectedCheckFriendship{false, nil},
			expectedAddFriendship:    nil,
			expectedErr:              nil,
		},
		{
			Name:                     "User Email not found",
			Input:                    []string{"firstuser@example.com", "seconduser@example.com"},
			expectedCheckUserByEmail: expectedCheckUserByEmail{false, nil},
			expectedCheckFriendship:  expectedCheckFriendship{false, nil},
			expectedAddFriendship:    nil,
			expectedErr:              UserNotFound,
		},
		{
			Name:                     "Server error from CheckUserByEmail",
			Input:                    []string{"firstuser@example.com", "seconduser@example.com"},
			expectedCheckUserByEmail: expectedCheckUserByEmail{true, ServerError},
			expectedCheckFriendship:  expectedCheckFriendship{false, nil},
			expectedAddFriendship:    nil,
			expectedErr:              ServerError,
		},
		{
			Name:                     "Friendship already added",
			Input:                    []string{"firstuser@example.com", "seconduser@example.com"},
			expectedCheckUserByEmail: expectedCheckUserByEmail{true, nil},
			expectedCheckFriendship:  expectedCheckFriendship{true, nil},
			expectedAddFriendship:    nil,
			expectedErr:              FriendshipExisted,
		},
		{Name: "Server error from CheckFriendship",
			Input:                    []string{"firstuser@example.com", "seconduser@example.com"},
			expectedCheckUserByEmail: expectedCheckUserByEmail{true, nil},
			expectedCheckFriendship:  expectedCheckFriendship{false, ServerError},
			expectedAddFriendship:    nil,
			expectedErr:              ServerError,
		},
		{
			Name:                     "Server error from AddFriendship",
			Input:                    []string{"firstuser@example.com", "seconduser@example.com"},
			expectedCheckUserByEmail: expectedCheckUserByEmail{true, nil},
			expectedCheckFriendship:  expectedCheckFriendship{false, nil},
			expectedAddFriendship:    ServerError,
			expectedErr:              ServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Setup Instance
			repo := &repository.MockRepository{}
			ctrl := New(repo)
			ctx := context.Background()

			// Defined mock Behaviors
			for _, inputcase := range tc.Input {
				repo.On("CheckUserByEmail", ctx, inputcase).
					Return(tc.expectedCheckUserByEmail.expectedExist, tc.expectedCheckUserByEmail.expectedErr)
			}
			repo.On("CheckFriendship", ctx, tc.Input).
				Return(tc.expectedCheckFriendship.expectedExist, tc.expectedCheckFriendship.expectedErr)

			repo.On("AddFriendship", ctx, tc.Input[0], tc.Input[1]).
				Return(tc.expectedAddFriendship)

			// Run the Test
			err := ctrl.AddFriends(ctx, tc.Input)

			// Check Result
			if err != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, tc.expectedErr)
			}

		})
	}
}
