package repository

import (
	"assignment/internal/repository/testdata"
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestRepoImplement_CheckFriendship(t *testing.T) {
	testCases := []struct {
		Name          string
		UserEmails    []string
		ExpectedExist bool
		ExpectedErr   error
		WantDBFail    bool
	}{
		{
			Name:          "Success",
			UserEmails:    []string{"new_email_1@example.com", "new_email_2@example.com"},
			ExpectedExist: false,
			ExpectedErr:   nil,
			WantDBFail:    false,
		},
		{
			Name:          "Friendship already added",
			UserEmails:    []string{"already_exist_email_1@example.com", "already_exist_email_2@example.com"},
			ExpectedExist: true,
			ExpectedErr:   nil,
			WantDBFail:    false,
		},
		{
			Name:          "Internal Server Error",
			WantDBFail:    true,
			UserEmails:    []string{"internal-error-email_1@example.com", "internal-error-email_2@example.com"},
			ExpectedExist: false,
			ExpectedErr:   InternalErrorCheckFriendship,
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

				exists, err := repo.CheckFriendship(ctx, tc.UserEmails)

				if err != nil {
					require.EqualError(t, err, tc.ExpectedErr.Error())
				} else {
					require.Equal(t, tc.ExpectedExist, exists)
				}
			})
		})
	}

}
