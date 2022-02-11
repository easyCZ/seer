package main

import (
	"context"
	"log"
	"net/http"

	"github.com/easyCZ/qfy/internal/api"
	"github.com/easyCZ/qfy/internal/db"
	"github.com/go-chi/chi/v5"

	"github.com/spf13/pflag"
)

func main() {

	var (
		dbHost     = pflag.String("db-host", "localhost", "Database hostname")
		dbPort     = pflag.Int("db-port", 5432, "Database port")
		dbUser     = pflag.String("db-user", "gitpod", "Database user")
		dbPassword = pflag.String("db-password", "gitpod", "Database user password")
		dbName     = pflag.String("db-name", "postgres", "Database name")
	)

	pflag.Parse()

	database, err := db.New(db.ConnectionParams{
		Host:         *dbHost,
		Port:         *dbPort,
		User:         *dbUser,
		Password:     *dbPassword,
		DatabaseName: *dbName,
	})
	if err != nil {
		log.Fatalf("Failed to setup db: %v", err)
	}

	if err := db.Migrate(database); err != nil {
		log.Fatalf("Failed to migrate DB: %v", err)
	}

	repo := db.NewSyntheticsRepository(database)
	service := api.NewSyntheticsService(repo)

	if _, err := repo.Create(context.Background(), &db.Synthetic{
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
	}); err != nil {
		log.Fatalf("Failed to create synthetic: %v", err)
	}

	r := chi.NewRouter()
	api.AddRoutes(r, service)

	log.Printf("Starting server on http://localhost:3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
