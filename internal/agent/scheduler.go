package agent

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	apiv1 "github.com/easyCZ/seer/gen/v1"
	"k8s.io/apimachinery/pkg/util/rand"
)

// Scheduler schedules Runners to execute a Synthetic according to specification.
// Scheduler should be used a singleton. The scheduler does not store store results of invocations.
type Scheduler struct {
	synthetics []*apiv1.Synthetic

	runClient apiv1.RunServiceClient
}

func (s *Scheduler) SingleRun(ctx context.Context) error {
	var wg sync.WaitGroup
	for _, synthetic := range s.synthetics {
		wg.Add(1)
		go (func(syn *apiv1.Synthetic) {
			defer wg.Done()
			_, _ = s.executeRunner(ctx, synthetic)
		})(synthetic)
	}

	wg.Wait()
	return nil
}

func (s *Scheduler) executeRunner(ctx context.Context, syn *apiv1.Synthetic) (*apiv1.Run, error) {
	r := &Runner{
		ID:     "runner-id",
		Client: http.DefaultClient,
		debug:  true,
	}

	result, err := r.Execute(ctx, syn)
	if err != nil {
		return nil, fmt.Errorf("runner failed: %w", err)
	}

	created, err := s.runClient.Create(ctx, &apiv1.CreateRequest{
		Run: &apiv1.Run{
			Id:          rand.String(5),
			SyntheticId: syn.Id,
			Results:     result,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to store run in hub service: %w", err)
	}

	return created.Run, nil
}
