package testdata

import (
	"context"
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

var appDB *sql.DB

func TestDatabase(t *testing.T, testScenario func(tx *sql.Tx)) {

	if appDB == nil {
		var err error
		appDB, err = sql.Open("postgres", "postgresql://test:@@localhost:5432/test?sslmode=disable")
		require.NoError(t, err)
	}

	tx, err := appDB.BeginTx(context.Background(), nil)
	require.NoError(t, err)

	defer tx.Rollback()

	testScenario(tx)
}

func LoadTestSQLFile(t *testing.T, db *sql.Tx, file string) {
	script, err := os.ReadFile(file)
	require.NoError(t, err)

	_, err = db.Exec(string(script))
	require.NoError(t, err)
}
