package srv

import (
	"fmt"
	"net"

	apiv1 "github.com/easyCZ/seer/gen/v1"
	"github.com/easyCZ/seer/internal/db"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type CPConfig struct {
	Logger *zap.SugaredLogger

	DB       db.ConnectionParams
	GRPCPort int
}

func ListenAndServeControlPlane(c CPConfig) error {
	logger := c.Logger

	database, err := setupDB(c.DB)
	if err != nil {
		return fmt.Errorf("failed to setup db: %w", err)
	}

	syntheticsRepo := db.NewSyntheticsRepository(database)
	agentsRepo := db.NewAgentsRepository(database)

	syntheticsSvc := &SyntheticsService{repo: syntheticsRepo}
	agentsSvc := &AgentService{repo: agentsRepo}
	runSvc := &RunService{}

	grpcServer := grpc.NewServer()

	apiv1.RegisterAgentServiceServer(grpcServer, agentsSvc)
	apiv1.RegisterSyntheticsServiceServer(grpcServer, syntheticsSvc)
	apiv1.RegisterRunServiceServer(grpcServer, runSvc)

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", c.GRPCPort))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	logger.Infof("Starting gRPC server on localhost:%d", c.GRPCPort)
	if err := grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC: %w", err)
	}

	logger.Info("Finished serving gRPC API.")
	return nil
}

func setupDB(params db.ConnectionParams) (*gorm.DB, error) {
	database, err := db.New(params)
	if err != nil {
		return nil, fmt.Errorf("failed to setup db: %v", err)
	}

	if err := db.Migrate(database); err != nil {
		return nil, fmt.Errorf("failed to migrate DB: %v", err)
	}

	return database, nil
}
