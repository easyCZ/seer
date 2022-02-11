package agent

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {
	require.NoError(t, RunClient())
}
