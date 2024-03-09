package controller

import (
	"assignment/internal/repository"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMockController_BlockUsers(t *testing.T) {
	type expectedCheckUserByEmail struct {
		expectedExist bool
		expectedErr   error
	}
	type expectedCheckIfFollowed struct {
		expectedExist bool
		expectedErr   error
	}
	type expectedCheckIfBlocked struct {
		expectedExist bool
		expectedErr   error
	}

	testCases := []struct {
		Name  string
		Input []string
		expectedCheckUserByEmail
		expectedCheckIfBlocked
		expectedCheckIfFollowed
		expectedBlockToSubscribed    error
		expectedBlockToNonSubscribed error
		expectedErr                  error
	}{
		{
			Name:                         "Success",
			Input:                        []string{"requester@example.com", "target@example.com"},
			expectedCheckUserByEmail:     expectedCheckUserByEmail{true, nil},
			expectedCheckIfBlocked:       expectedCheckIfBlocked{false, nil},
			expectedCheckIfFollowed:      expectedCheckIfFollowed{false, nil},
			expectedBlockToNonSubscribed: nil,
			expectedBlockToSubscribed:    nil,
			expectedErr:                  nil,
		},
		{
			Name:                         "User Email not found",
			Input:                        []string{"requester@example.com", "target@example.com"},
			expectedCheckUserByEmail:     expectedCheckUserByEmail{false, nil},
			expectedCheckIfFollowed:      expectedCheckIfFollowed{false, nil},
			expectedCheckIfBlocked:       expectedCheckIfBlocked{false, nil},
			expectedBlockToNonSubscribed: nil,
			expectedBlockToSubscribed:    nil,
			expectedErr:                  UserNotFound,
		},
		{
			Name:                         "Server error from CheckUserByEmail",
			Input:                        []string{"requester@example.com", "target@example.com"},
			expectedCheckUserByEmail:     expectedCheckUserByEmail{true, ServerError},
			expectedCheckIfBlocked:       expectedCheckIfBlocked{false, nil},
			expectedCheckIfFollowed:      expectedCheckIfFollowed{false, nil},
			expectedBlockToNonSubscribed: nil,
			expectedBlockToSubscribed:    nil,
			expectedErr:                  ServerError,
		},
		{
			Name:                         "Already Blocked",
			Input:                        []string{"requester@example.com", "target@example.com"},
			expectedCheckUserByEmail:     expectedCheckUserByEmail{true, nil},
			expectedCheckIfBlocked:       expectedCheckIfBlocked{true, nil},
			expectedCheckIfFollowed:      expectedCheckIfFollowed{false, nil},
			expectedBlockToNonSubscribed: nil,
			expectedBlockToSubscribed:    nil,
			expectedErr:                  AlreadyBlocked,
		},
		{
			Name:                         "Server error from CheckIfBlocked",
			Input:                        []string{"requester@example.com", "target@example.com"},
			expectedCheckUserByEmail:     expectedCheckUserByEmail{true, nil},
			expectedCheckIfBlocked:       expectedCheckIfBlocked{false, ServerError},
			expectedCheckIfFollowed:      expectedCheckIfFollowed{false, nil},
			expectedBlockToNonSubscribed: nil,
			expectedBlockToSubscribed:    nil,
			expectedErr:                  ServerError,
		},
		{
			Name:                         "Server error from CheckSubscription",
			Input:                        []string{"requester@example.com", "target@example.com"},
			expectedCheckUserByEmail:     expectedCheckUserByEmail{true, nil},
			expectedCheckIfBlocked:       expectedCheckIfBlocked{false, nil},
			expectedCheckIfFollowed:      expectedCheckIfFollowed{false, ServerError},
			expectedBlockToNonSubscribed: nil,
			expectedBlockToSubscribed:    nil,
			expectedErr:                  ServerError,
		},
		{
			Name:                         "Server error from BlockToSubscribed",
			Input:                        []string{"requester@example.com", "target@example.com"},
			expectedCheckUserByEmail:     expectedCheckUserByEmail{true, nil},
			expectedCheckIfBlocked:       expectedCheckIfBlocked{false, nil},
			expectedCheckIfFollowed:      expectedCheckIfFollowed{true, nil},
			expectedBlockToNonSubscribed: nil,
			expectedBlockToSubscribed:    ServerError,
			expectedErr:                  ServerError,
		},
		{
			Name:                         "Server error from BlockToNonSubscribed",
			Input:                        []string{"requester@example.com", "target@example.com"},
			expectedCheckUserByEmail:     expectedCheckUserByEmail{true, nil},
			expectedCheckIfBlocked:       expectedCheckIfBlocked{false, nil},
			expectedCheckIfFollowed:      expectedCheckIfFollowed{false, nil},
			expectedBlockToNonSubscribed: ServerError,
			expectedBlockToSubscribed:    nil,
			expectedErr:                  ServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			repo := &repository.MockRepository{}
			ctrl := New(repo)
			ctx := context.Background()

			for _, inputCase := range tc.Input {
				repo.On("CheckUserByEmail", ctx, inputCase).
					Return(tc.expectedCheckUserByEmail.expectedExist, tc.expectedCheckUserByEmail.expectedErr)
			}
			repo.On("CheckIfFollowed", ctx, tc.Input).
				Return(tc.expectedCheckIfFollowed.expectedExist, tc.expectedCheckIfFollowed.expectedErr)
			repo.On("CheckIfBlocked", ctx, tc.Input).
				Return(tc.expectedCheckIfBlocked.expectedExist, tc.expectedCheckIfBlocked.expectedErr)
			repo.On("BlockToSubscribed", ctx, tc.Input).
				Return(tc.expectedBlockToSubscribed)
			repo.On("BlockToNonSubscribed", ctx, tc.Input).
				Return(tc.expectedBlockToNonSubscribed)

			err := ctrl.BlockUsers(ctx, tc.Input)
			if err != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, tc.expectedErr)
			}
		})
	}
}
