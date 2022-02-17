package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAgentUpsert(t *testing.T) {
	db := NewTestDB(t)
	repo := NewAgentsRepository(db)

	id := "test-agent-1"
	toUpsert := &Agent{
		ID:        id,
		Region:    "EU",
		IP:        "0.0.0.0",
		Connected: true,
	}

	_, err := repo.Upsert(context.Background(), toUpsert)
	require.NoError(t, err)

	get, err := repo.Get(context.Background(), id)
	require.NoError(t, err)

	require.Equal(t, toUpsert, get)
}
