package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAgentUpsert(t *testing.T) {
	db := NewTestDB(t)
	repo := NewAgentsRepository(db)

	toCreate := &Agent{
		Region:    "EU",
		IP:        "0.0.0.0",
		Connected: true,
	}

	created, err := repo.Create(context.Background(), toCreate)
	require.NoError(t, err)
	require.True(t, created.Connected)

	_, err = repo.Get(context.Background(), created.ID)
	require.NoError(t, err)

	require.NoError(t, repo.SetConnected(context.Background(), created.ID, false))

	get, err := repo.Get(context.Background(), created.ID)
	require.NoError(t, err)
	require.False(t, get.Connected)
}
