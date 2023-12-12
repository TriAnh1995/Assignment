package repository

import (
	"assignment/internal/repository/testdata"
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestRepoImplement_GetFollowedList(t *testing.T) {
	testCases := []struct {
		Name           string
		UserEmail      string
		ExpectedResult []string
		ExpectedErr    error
		WantDBFail     bool
	}{
		{
			Name:           "Success",
			UserEmail:      "target1@example.com",
			ExpectedResult: []string{"requester1@example.com"},
			ExpectedErr:    nil,
			WantDBFail:     false,
		},
		{
			Name:           "Internal Server Error",
			UserEmail:      "target2@example.com",
			ExpectedResult: nil,
			ExpectedErr:    InternalErrorGetFollowedList,
			WantDBFail:     true,
		},
	}
	ctx := context.Background()
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testdata.TestDatabase(t, func(tx *sql.Tx) {
				testdata.LoadTestSQLFile(t, tx, "testdata/testdata_for_subscription.sql")
				repo := New(tx)
				if tc.WantDBFail {
					dbMock, _, _ := sqlmock.New()
					repo = New(dbMock)
				}

				listOfFollowed, err := repo.GetFollowedList(ctx, tc.UserEmail)
				if err == nil {
					require.Equal(t, listOfFollowed, tc.ExpectedResult)
				} else {
					require.EqualError(t, err, tc.ExpectedErr.Error())
				}
			})
		})
	}
}
