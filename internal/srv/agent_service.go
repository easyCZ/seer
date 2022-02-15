package srv

import (
	"fmt"
	"github.com/easyCZ/qfy/gen/v1"
	"log"
	"time"
)

type AgentService struct {
}

func (s *AgentService) Subscribe(r *agentv1.SubscribeRequest, stream agentv1.AgentService_SubscribeServer) error {
	log.Printf("Received subscription: %v", r)
	t := time.NewTicker(5 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			if err := stream.Send(&agentv1.SubscribeResponse{}); err != nil {
				return fmt.Errorf("failed to send response to client: %w", err)
			}
		case <-stream.Context().Done():
			return nil
		}
	}
}
