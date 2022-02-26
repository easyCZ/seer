package integration

import (
	"testing"

	"github.com/easyCZ/seer/internal/db"
	"github.com/easyCZ/seer/internal/hub"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func Test(t *testing.T) {
	logger := zaptest.NewLogger(t)
	srv, err := hub.NewServer(logger.Sugar(), hub.Config{
		DB:       db.TestDBConnectionParams(),
		GRPCPort: 4000,
	})
	require.NoError(t, err)

	startHub(t, srv)
}

func startHub(t *testing.T, srv *hub.Server) {
	t.Helper()

	t.Cleanup(func() {
		srv.GracefulStop()
	})

	go (func() {
		require.NoError(t, srv.ListenAndServe(), "must start grpc server")
	})()
}
