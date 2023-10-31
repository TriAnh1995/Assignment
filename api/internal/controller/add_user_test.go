package controller

import (
	"assignment/internal/model"
	"assignment/internal/repository"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
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
			expectedErr:               UserAlreadyExist,
		},
		{
			Name:                      "Server error from CheckUserByEmail",
			input:                     model.User{Name: "Normal_User", Email: "normal_user@example.com"},
			expectedCheckUserByEmail:  expectedCheckUserByEmail{false, ServerError},
			expectedAddUserToDatabase: nil,
			expectedErr:               ServerError,
		},
		{
			Name:                      "Server error from AddUserToDatabase",
			input:                     model.User{Name: "Normal_User", Email: "normal_user@example.com"},
			expectedCheckUserByEmail:  expectedCheckUserByEmail{false, nil},
			expectedAddUserToDatabase: ServerError,
			expectedErr:               ServerError,
		},
	}

	// Setup Instance
	repo := &repository.MockRepository{}
	ctrl := New(repo)
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

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
