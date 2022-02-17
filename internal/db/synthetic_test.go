package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSyntheticCRUD(t *testing.T) {
	db := NewTestDB(t)

	repo := NewSyntheticsRepository(db)
	created, err := repo.Create(context.Background(), &Synthetic{
		ID:   1,
		Name: "my synthetic",
		Spec: SyntheticSpec{},
	})
	require.NoError(t, err)

	fmt.Println(created)
}
