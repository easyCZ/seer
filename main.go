package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/easyCZ/qfy/internal/api"
	"github.com/easyCZ/qfy/internal/db"
	"github.com/go-chi/chi/v5"
)

func main() {
	database, err := db.New(db.ConnectionParams{
		Host:         "localhost",
		Port:         38386,
		User:         "gitpod",
		Password:     "gitpod",
		DatabaseName: "postgres",
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

		enc := json.NewEncoder(rw)
		if err := enc.Encode(syns); err != nil {
			log.Println("Failed to serialize synthetics to response")
			return
		}

		return
	})

	log.Printf("Starting server on http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
