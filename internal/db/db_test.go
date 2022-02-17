package db

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// NewTestDB creates a new DB instance for tests.
// You may need to create the test database using:
// 	createdb -p 5432 -h 127.0.0.1 -U gitpod -e postgres-test
func NewTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	database, err := New(ConnectionParams{
		Host:         "localhost",
		Port:         5432,
		User:         "gitpod",
		Password:     "gitpod",
		DatabaseName: "postgres-test",
	})
	require.NoError(t, err)
	require.NoError(t, Migrate(database))

	t.Cleanup(func() {
		require.NoError(t, DropTables(database))
	})

	return database
}
