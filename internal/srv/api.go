package srv

import (
	"github.com/easyCZ/qfy/internal/api"
	"github.com/easyCZ/qfy/internal/db"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type CPConfig struct {
	DB db.ConnectionParams
}

func ListenAndServeControlPlane(c CPConfig) error {
	database, err := db.New(c.DB)
	if err != nil {
		log.Fatalf("Failed to setup db: %v", err)
	}

	if err := db.Migrate(database); err != nil {
		log.Fatalf("Failed to migrate DB: %v", err)
	}

	repo := db.NewSyntheticsRepository(database)
	service := api.NewSyntheticsService(repo)

	r := chi.NewRouter()
	api.AddRoutes(r, service)

	log.Printf("Starting server on http://localhost:3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}

	log.Printf("Finished serving API.")
	return nil
}
