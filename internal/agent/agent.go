package agent

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	apiv1 "github.com/easyCZ/qfy/gen/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Agent struct {
	URL      string
	TickRate time.Duration

	synthetics   []*apiv1.Synthetic
	syntheticsMu sync.RWMutex

	wg sync.WaitGroup
}

func (a *Agent) Run() {
	a.wg.Add(1)
	go (func() {
		defer a.wg.Done()

		ctx := context.Background()
		var client apiv1.SyntheticsServiceClient

		ticker := time.NewTicker(time.Second * 5)

		for {
			select {
			case <-ticker.C:
				log.Printf("Tick")
				var clientErr error
				if client == nil {
					client, clientErr = a.newSyntheticsClient()
					if clientErr != nil {
						log.Printf("Failed to setup synthetics client: %v", clientErr)
						continue
					}
				}

				synthetics, err := a.listSynthetics(ctx, client)
				if err != nil {
					log.Printf("Failed to list synthetics: %v", err)
					continue
				}

				log.Printf("Received %d synthetics from upstream", len(synthetics))
				a.syntheticsMu.Lock()
				a.synthetics = synthetics
				a.syntheticsMu.Unlock()

			}
		}
	})()

	a.wg.Add(1)
	go (func() {
		defer a.wg.Done()
		ctx := context.Background()

		// for now, run a synthetic for each tick
		for {
			ticker := time.NewTicker(time.Second * 10)

			select {
			case <- ticker.C:
				a.syntheticsMu.RLock()
				
			}
		}
		
		for _, s := range a.synthetics
		ExecuteSynthetic(ctx, )
	})()

	a.wg.Wait()

}

func (a *Agent) newSyntheticsClient() (apiv1.SyntheticsServiceClient, error) {
	conn, err := grpc.Dial(a.URL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to establish connection to control plane: %w", err)
	}

	return apiv1.NewSyntheticsServiceClient(conn), nil
}

func (a *Agent) listSynthetics(ctx context.Context, client apiv1.SyntheticsServiceClient) ([]*apiv1.Synthetic, error) {
	stream, err := client.ListSynthetics(ctx, &apiv1.ListSyntheticsRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed to establish stream of synthetics: %w", err)
	}

	var synthetics []*apiv1.Synthetic
	for {
		synthetic, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return synthetics, nil
			}

			return nil, fmt.Errorf("error receiving synthetic from stream: %w", err)
		}

		synthetics = append(synthetics, synthetic)
	}
}
