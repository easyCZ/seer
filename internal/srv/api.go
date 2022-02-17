package srv

import (
	agentv1 "github.com/easyCZ/qfy/gen/v1"
	"github.com/easyCZ/qfy/internal/api"
	"github.com/easyCZ/qfy/internal/db"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"log"
	"net"
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

	agentsRepo := db.NewAgentsRepository(database)

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

	listener, err := net.Listen("tcp", "localhost:3001")
	if err != nil {
		log.Fatalf("Failed to listen on port 3001, %v", err)
	}

	grpcServer := grpc.NewServer()
	agentv1.RegisterAgentServiceServer(grpcServer, &AgentService{
		repo: agentsRepo,
	})

	go func() {
		log.Printf("Starting gRPC server on localhost:3001")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("gRPC server failed to start: %v", err)
		}
	}()

	r := chi.NewRouter()
	api.AddRoutes(r, service)

	log.Printf("Starting server on http://localhost:3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}

	log.Printf("Finished serving API.")
	return nil
}
