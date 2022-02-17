package srv

import (
	"context"
	"fmt"
	"log"
	"time"

	agentv1 "github.com/easyCZ/qfy/gen/v1"
	"github.com/easyCZ/qfy/internal/db"
	"google.golang.org/grpc/peer"
)

type AgentService struct {
	repo *db.AgentsRepository
}

func (s *AgentService) Subscribe(r *agentv1.SubscribeRequest, stream agentv1.AgentService_SubscribeServer) error {
	log.Printf("Received subscription from client%s", r.AgentID)

	p, _ := peer.FromContext(stream.Context())

	agent, err := s.repo.Create(stream.Context(), &db.Agent{
		Region:    r.Location,
		IP:        p.Addr.String(),
		Connected: true,
	})
	if err != nil {
		return fmt.Errorf("failed to register agent: %v", err)
	}

	t := time.NewTicker(5 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			if err := stream.Send(&agentv1.SubscribeResponse{}); err != nil {
				return fmt.Errorf("failed to send response to client: %w", err)
			}
		case <-stream.Context().Done():
			log.Printf("client %s disconnected", r.AgentID)

			s.repo.SetConnected(context.Background(), agent.ID, false)
			return nil
		}
	}
}
