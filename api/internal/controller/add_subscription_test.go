package controller

import (
	"assignment/internal/repository"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestControllerImplement_AddSubscription(t *testing.T) {
	type expectedCheckUserByEmail struct {
		expectedExist bool
		expectedErr   error
	}
	type expectedCheckIfSubscribed struct {
		expectedExist bool
		expectedErr   error
	}
	type expectedCheckBlocked struct {
		expectedExist bool
		expectedErr   error
	}

	testCases := []struct {
		Name  string
		Input []string
		expectedCheckUserByEmail
		expectedCheckIfSubscribed
		expectedCheckBlocked
		expectedSubscribeToBlocked    error
		expectedSubscribeToNonBlocked error
		expectedErr                   error
	}{
		{
			Name:                          "Success",
			Input:                         []string{"requester@example.com", "target@example.com"},
			expectedCheckUserByEmail:      expectedCheckUserByEmail{true, nil},
			expectedCheckIfSubscribed:     expectedCheckIfSubscribed{false, nil},
			expectedCheckBlocked:          expectedCheckBlocked{false, nil},
			expectedSubscribeToBlocked:    nil,
			expectedSubscribeToNonBlocked: nil,
			expectedErr:                   nil,
		},
		{
			Name:                          "User Email not found",
			Input:                         []string{"requester@example.com", "target@example.com"},
			expectedCheckUserByEmail:      expectedCheckUserByEmail{false, nil},
			expectedCheckIfSubscribed:     expectedCheckIfSubscribed{false, nil},
			expectedCheckBlocked:          expectedCheckBlocked{false, nil},
			expectedSubscribeToBlocked:    nil,
			expectedSubscribeToNonBlocked: nil,
			expectedErr:                   UserNotFound,
		},
		{
			Name:                          "Server error from CheckUserByEmail",
			Input:                         []string{"requester@example.com", "target@example.com"},
			expectedCheckUserByEmail:      expectedCheckUserByEmail{true, ServerError},
			expectedCheckIfSubscribed:     expectedCheckIfSubscribed{false, nil},
			expectedCheckBlocked:          expectedCheckBlocked{false, nil},
			expectedSubscribeToBlocked:    nil,
			expectedSubscribeToNonBlocked: nil,
			expectedErr:                   ServerError,
		},
		{
			Name:                          "Already Subscribed",
			Input:                         []string{"requester@example.com", "target@example.com"},
			expectedCheckUserByEmail:      expectedCheckUserByEmail{true, nil},
			expectedCheckIfSubscribed:     expectedCheckIfSubscribed{true, nil},
			expectedCheckBlocked:          expectedCheckBlocked{false, nil},
			expectedSubscribeToBlocked:    nil,
			expectedSubscribeToNonBlocked: nil,
			expectedErr:                   AlreadySubscribed,
		},
		{
			Name:                          "Server error from CheckSubscription",
			Input:                         []string{"requester@example.com", "target@example.com"},
			expectedCheckUserByEmail:      expectedCheckUserByEmail{true, nil},
			expectedCheckIfSubscribed:     expectedCheckIfSubscribed{false, ServerError},
			expectedCheckBlocked:          expectedCheckBlocked{false, nil},
			expectedSubscribeToBlocked:    nil,
			expectedSubscribeToNonBlocked: nil,
			expectedErr:                   ServerError,
		},
		{
			Name:                          "Server error from CheckBlocked",
			Input:                         []string{"requester@example.com", "target@example.com"},
			expectedCheckUserByEmail:      expectedCheckUserByEmail{true, nil},
			expectedCheckIfSubscribed:     expectedCheckIfSubscribed{false, nil},
			expectedCheckBlocked:          expectedCheckBlocked{false, ServerError},
			expectedSubscribeToBlocked:    nil,
			expectedSubscribeToNonBlocked: nil,
			expectedErr:                   ServerError,
		},
		{
			Name:                          "Server error from SubscribeToBlocked",
			Input:                         []string{"requester@example.com", "target@example.com"},
			expectedCheckUserByEmail:      expectedCheckUserByEmail{true, nil},
			expectedCheckIfSubscribed:     expectedCheckIfSubscribed{false, nil},
			expectedCheckBlocked:          expectedCheckBlocked{true, nil},
			expectedSubscribeToBlocked:    ServerError,
			expectedSubscribeToNonBlocked: nil,
			expectedErr:                   ServerError,
		},
		{
			Name:                          "Server error from SubscribeToNonBlocked",
			Input:                         []string{"requester@example.com", "target@example.com"},
			expectedCheckUserByEmail:      expectedCheckUserByEmail{true, nil},
			expectedCheckIfSubscribed:     expectedCheckIfSubscribed{false, nil},
			expectedCheckBlocked:          expectedCheckBlocked{false, nil},
			expectedSubscribeToBlocked:    nil,
			expectedSubscribeToNonBlocked: ServerError,
			expectedErr:                   ServerError,
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
			repo.On("CheckIfSubscribed", ctx, tc.Input).
				Return(tc.expectedCheckIfSubscribed.expectedExist, tc.expectedCheckIfSubscribed.expectedErr)
			repo.On("CheckIfBlocked", ctx, tc.Input).
				Return(tc.expectedCheckBlocked.expectedExist, tc.expectedCheckBlocked.expectedErr)
			repo.On("SubscribeToBlocked", ctx, tc.Input).
				Return(tc.expectedSubscribeToBlocked)
			repo.On("SubscribeToNonBlocked", ctx, tc.Input).
				Return(tc.expectedSubscribeToNonBlocked)

			err := ctrl.AddSubscription(ctx, tc.Input)
			if err != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, tc.expectedErr)
			}
		})
	}
}
