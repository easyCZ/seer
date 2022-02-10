package main

import (
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

	if err := database.AutoMigrate(&db.Synthetic{}); err != nil {
		log.Fatalf("Failed to migrate DB: %v", err)
	}

	result := database.Create(&db.Synthetic{
		Name: "my-first-synthetic",
	})
	if result.Error != nil {
		log.Fatalf("Failed to create db record: %v", err)
	}

	repo := db.NewSyntheticsRepository(database)
	service := api.NewSyntheticsService(repo)

	r := chi.NewRouter()
	r.Get("/synthetics", func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		syns, err := service.List(ctx)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
		}

		if err := api.WriteJSON(rw, r, http.StatusOK, syns); err != nil {
			log.Println("Failed to serialize synthetics to response")
		}
	})

	log.Printf("Starting server on http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
