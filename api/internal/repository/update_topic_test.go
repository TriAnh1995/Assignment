package repository

import (
	"assignment/internal/model"
	"assignment/internal/repository/testdata"
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestRepoImplement_UpdateTopic(t *testing.T) {
	testCases := []struct {
		Name        string
		User        model.UpdateInfo
		ExpectedErr error
		WantDBFail  bool
	}{
		{
			Name: "Success",
			User: model.UpdateInfo{
				Sender: "update_user_email@example.com",
				Text:   "Hello World! mentioned_email@example.com",
			},
			ExpectedErr: nil,
			WantDBFail:  false,
		},
		{
			Name: "Internal Server Error",
			User: model.UpdateInfo{
				Sender: "update_user_email@example.com",
				Text:   "Hello World! mentioned_email@example.com",
			},
			ExpectedErr: InternalErrorUpdateTopic,
			WantDBFail:  true,
		},
	}
	ctx := context.Background()
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testdata.TestDatabase(t, func(tx *sql.Tx) {
				testdata.LoadTestSQLFile(t, tx, "testdata/testdata_for_user_accounts.sql")
				repo := New(tx)
				if tc.WantDBFail {
					dbMock, _, _ := sqlmock.New()
					repo = New(dbMock)
				}
				err := repo.UpdateTopic(ctx, tc.User)
				if err != nil {
					require.EqualError(t, err, tc.ExpectedErr.Error())
				} else {
					require.Equal(t, err, tc.ExpectedErr)
				}
			})
		})
	}
}
