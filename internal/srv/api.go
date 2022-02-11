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

	//if _, err := repo.Create(context.Background(), &db.Synthetic{
	//	Name: "initial-synthetic",
	//	Spec: db.SyntheticSpec{
	//		Variables: []*db.Variable{
	//			{
	//				Name:  "TEST",
	//				Value: "foo",
	//			},
	//		},
	//		Steps: []*db.StepSpec{
	//			{
	//				URL: "https://jsonplaceholder.typicode.com/posts",
	//				Headers: map[string]string{
	//					"Content-Type": "application/json",
	//				},
	//			},
	//		},
	//	},
	//}); err != nil {
	//	log.Fatalf("Failed to create synthetic: %v", err)
	//}

	r := chi.NewRouter()
	api.AddRoutes(r, service)

	log.Printf("Starting server on http://localhost:3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}

	log.Printf("Finished serving API.")
	return nil
}
