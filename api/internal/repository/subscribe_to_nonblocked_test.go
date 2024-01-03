package repository

import (
	"assignment/internal/repository/testdata"
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestRepoImplement_SubscribeToNonBlocked(t *testing.T) {
	testCases := []struct {
		Name        string
		UserEmails  []string
		ExpectedErr error
		WantDBFail  bool
	}{
		{
			Name:        "Success",
			UserEmails:  []string{"requester@example.com", "target@example.com"},
			ExpectedErr: nil,
			WantDBFail:  false,
		},
		{
			Name:        "Internal Server Error",
			UserEmails:  []string{"requester@example.com", "target@example.com"},
			ExpectedErr: InternalErrorSubscribeToNonBlocked,
			WantDBFail:  true,
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
				err := repo.SubscribeToNonBlocked(ctx, tc.UserEmails)
				if err != nil {
					require.EqualError(t, err, tc.ExpectedErr.Error())
				} else {
					require.Equal(t, err, tc.ExpectedErr)
				}
			})
		})
	}
}
