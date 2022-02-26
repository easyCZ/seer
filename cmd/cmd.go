package main

import (
	"context"
	"fmt"

	agentv1 "github.com/easyCZ/seer/gen/v1"
	"github.com/easyCZ/seer/internal/db"
	"github.com/easyCZ/seer/internal/log"
	"github.com/easyCZ/seer/internal/srv"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/apimachinery/pkg/util/rand"
)

var (
	rootCmd = &cobra.Command{
		Use:   "qfy",
		Short: "Qfy - Self-hosted synthetics",
	}

	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start API server & Control Plane",
		Run: func(cmd *cobra.Command, args []string) {
			logger, _ := zap.NewDevelopment()
			if err := srv.ListenAndServeControlPlane(srv.CPConfig{
				Logger:   logger.Sugar(),
				GRPCPort: 3000,
				DB: db.ConnectionParams{
					Host:         dbHost,
					Port:         dbPort,
					User:         dbUser,
					Password:     dbPassword,
					DatabaseName: dbName,
				},
			}); err != nil {
				logger.Sugar().Fatalw("Failed to listen and serve", zap.Error(err))
			}
		},
	}

	agentCmd = &cobra.Command{
		Use:   "agent",
		Short: "Start agent",
		Run: func(cmd *cobra.Command, args []string) {
			if err := runClient(cmd.Context()); err != nil {
				log.FromContext(cmd.Context()).Fatalw("Failed to run agent.", zap.Error(err))
			}
		},
	}

	fixturesCmd = &cobra.Command{
		Use:   "fixtures",
		Short: "Generate database fixtures",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			repo, err := db.NewSyntheticsRepositoryFromDBParams(db.ConnectionParams{
				Host:         dbHost,
				Port:         dbPort,
				User:         dbUser,
				Password:     dbPassword,
				DatabaseName: dbName,
			})
			if err != nil {
				log.FromContext(ctx).Fatalw("Failed to create synthetics repository", zap.Error(err))
			}

			count := 10
			for i := 0; i < count; i++ {
				_, err := repo.Create(ctx, &db.Synthetic{
					Name: "initial-synthetic",
					Spec: db.SyntheticSpec{
						Variables: []*db.Variable{
							{
								Name:  "TEST",
								Value: "foo",
							},
						},
						Steps: []*db.StepSpec{
							{
								URL: "https://jsonplaceholder.typicode.com/posts",
								Headers: map[string]string{
									"Content-Type": "application/json",
								},
							},
						},
					},
				})
				if err != nil {
					log.FromContext(ctx).Fatalw("Failed to create record", zap.Error(err))
				}
			}

			log.FromContext(ctx).Infof("Created %d fixtures.", count)
		},
	}
)

func runClient(ctx context.Context) error {
	conn, err := grpc.Dial("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := agentv1.NewAgentServiceClient(conn)

	stream, err := client.Subscribe(ctx, &agentv1.SubscribeRequest{
		AgentID:  rand.String(10),
		Location: "EU",
	})
	if err != nil {
		return fmt.Errorf("failed to subscribe agent: %w", err)
	}

	logger := log.FromContext(ctx)
	for {
		msg, err := stream.Recv()
		if err != nil {
			return fmt.Errorf("failed to receive message: %w", err)
		}

		logger.Infof("Got message %v", msg)
	}

}

var (
	dbHost, dbUser, dbPassword, dbName string
	dbPort                             uint
)

func init() {
	serverCmd.Flags().StringVar(&dbHost, "db-host", "localhost", "Database hostname")
	serverCmd.Flags().UintVar(&dbPort, "db-port", 5432, "Database port")
	serverCmd.Flags().StringVar(&dbUser, "db-user", "gitpod", "Database user")
	serverCmd.Flags().StringVar(&dbPassword, "db-password", "gitpod", "Database user password")
	serverCmd.Flags().StringVar(&dbName, "db-name", "postgres", "Database name")

	fixturesCmd.Flags().StringVar(&dbHost, "db-host", "localhost", "Database hostname")
	fixturesCmd.Flags().UintVar(&dbPort, "db-port", 5432, "Database port")
	fixturesCmd.Flags().StringVar(&dbUser, "db-user", "gitpod", "Database user")
	fixturesCmd.Flags().StringVar(&dbPassword, "db-password", "gitpod", "Database user password")
	fixturesCmd.Flags().StringVar(&dbName, "db-name", "postgres", "Database name")

	rootCmd.AddCommand(serverCmd, agentCmd, fixturesCmd)
}

func main() {
	logger, _ := zap.NewDevelopment()
	sugared := logger.Sugar()
	ctx := log.ToContext(context.Background(), sugared)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		sugared.Fatal("Command execution failed.", zap.Error(err))
	}
}
