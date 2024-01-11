package controller

import (
	"assignment/internal/middleware"
	"assignment/internal/model"
	"assignment/internal/repository"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestImpl_AddUsers(t *testing.T) {
	type expectedCheckUserByEmail struct {
		expectedExist bool
		expectedErr   error
	}
	testCases := []struct {
		Name  string
		input model.User
		expectedCheckUserByEmail
		expectedAddUserToDatabase error
		expectedErr               error
	}{
		{
			Name:                      "Success",
			input:                     model.User{Name: "New_User", Email: "new_user@example.com"},
			expectedCheckUserByEmail:  expectedCheckUserByEmail{false, nil},
			expectedAddUserToDatabase: nil,
			expectedErr:               nil,
		},
		{
			Name:                      "Email Already Exist",
			input:                     model.User{Name: "Old_User", Email: "old_user@example.com"},
			expectedCheckUserByEmail:  expectedCheckUserByEmail{true, nil},
			expectedAddUserToDatabase: nil,
			expectedErr:               middleware.UserAlreadyExist,
		},
		{
			Name:                      "Server error from CheckUserByEmail",
			input:                     model.User{Name: "Normal_User", Email: "normal_user@example.com"},
			expectedCheckUserByEmail:  expectedCheckUserByEmail{false, middleware.ServerError},
			expectedAddUserToDatabase: nil,
			expectedErr:               middleware.ServerError,
		},
		{
			Name:                      "Server error from AddUserToDatabase",
			input:                     model.User{Name: "Normal_User", Email: "normal_user@example.com"},
			expectedCheckUserByEmail:  expectedCheckUserByEmail{false, nil},
			expectedAddUserToDatabase: middleware.ServerError,
			expectedErr:               middleware.ServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Setup Instance
			repo := &repository.MockRepository{}
			ctrl := New(repo)
			ctx := context.Background()
			// Defined mock Behaviors
			repo.On("CheckUserByEmail", ctx, tc.input.Email).
				Return(tc.expectedCheckUserByEmail.expectedExist, tc.expectedCheckUserByEmail.expectedErr)
			repo.On("AddUser", ctx, tc.input).
				Return(tc.expectedAddUserToDatabase)
			// Run the Test
			err := ctrl.AddUsers(ctx, tc.input)
			// Check Result
			if err != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, tc.expectedErr)
			}
		})
	}
}
