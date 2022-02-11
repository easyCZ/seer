package main

import (
	"context"
	"github.com/easyCZ/qfy/internal/db"
	"github.com/easyCZ/qfy/internal/srv"
	"github.com/spf13/cobra"
	"log"
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
			if err := srv.ListenAndServeControlPlane(srv.CPConfig{
				DB: db.ConnectionParams{
					Host:         dbHost,
					Port:         dbPort,
					User:         dbUser,
					Password:     dbPassword,
					DatabaseName: dbName,
				},
			}); err != nil {
				log.Fatalln("Failed to setup database connection.", err)
			}
		},
	}

	agentCmd = &cobra.Command{
		Use:   "agent",
		Short: "Start agent",
	}

	fixturesCmd = &cobra.Command{
		Use:   "fixtures",
		Short: "Generate database fixtures",
		Run: func(cmd *cobra.Command, args []string) {
			repo, err := db.NewSyntheticsRepositoryFromDBParams(db.ConnectionParams{
				Host:         dbHost,
				Port:         dbPort,
				User:         dbUser,
				Password:     dbPassword,
				DatabaseName: dbName,
			})
			if err != nil {
				log.Fatalln("Failed to create synthetics repository", err)
			}

			ctx := context.Background()
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
					log.Fatalln("Failed to create record", err)
				}
			}

			log.Printf("Created %d fixtures.", count)
		},
	}
)

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
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("Command execution failed.", err)
	}
}
