package repository

import (
	"assignment/internal/repository/testdata"
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestRepoImplement_GetFriendsList(t *testing.T) {
	testCases := []struct {
		Name           string
		UserEmail      string
		ExpectedResult []string
		ExpectedErr    error
		WantDBFail     bool
	}{
		{
			Name:           "Success",
			UserEmail:      "already_exist_email_1@example.com",
			ExpectedResult: []string{"already_exist_email_2@example.com", "already_exist_email_3@example.com"},
			ExpectedErr:    nil,
			WantDBFail:     false,
		},
		{
			Name:           "Empty friends list",
			UserEmail:      "zero_friends_email@example.com",
			ExpectedResult: nil,
			ExpectedErr:    nil,
			WantDBFail:     false,
		},
		{
			Name:           "Internal Server Error",
			UserEmail:      "already_exist_email_1@example.com",
			ExpectedResult: nil,
			ExpectedErr:    InternalErrorGetFriendList,
			WantDBFail:     true,
		},
	}
	ctx := context.Background()
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testdata.TestDatabase(t, func(tx *sql.Tx) {
				testdata.LoadTestSQLFile(t, tx, "testdata/testdata_for_friendship.sql")
				repo := New(tx)
				if tc.WantDBFail {
					dbMock, _, _ := sqlmock.New()
					repo = New(dbMock)
				}

				listOfFriends, err := repo.GetFriendsList(ctx, tc.UserEmail)

				if err == nil {
					require.Equal(t, listOfFriends, tc.ExpectedResult)
				} else {
					require.EqualError(t, err, tc.ExpectedErr.Error())
				}
			})
		})
	}
}
