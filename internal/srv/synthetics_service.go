package srv

import (
	"fmt"
	"net/http"

	apiv1 "github.com/easyCZ/qfy/gen/v1"
	"github.com/easyCZ/qfy/internal/db"
)

type SyntheticsService struct {
	repo *db.SyntheticsRepository
}

func (s *SyntheticsService) ListSynthetics(req *apiv1.ListSyntheticsRequest, stream apiv1.SyntheticsService_ListSyntheticsServer) error {
	synthetics, err := s.repo.List(stream.Context())
	if err != nil {
		return fmt.Errorf("failed to list synthetics: %w", err)
	}

	for _, s := range synthetics {
		if err := stream.Send(dbSyntheticToV1(s)); err != nil {
			return fmt.Errorf("failed to send synthetic to client: %w", err)
		}
	}

	return nil
}

func dbSyntheticToV1(s *db.Synthetic) *apiv1.Synthetic {
	return &apiv1.Synthetic{
		Id:   fmt.Sprintf("%d", s.ID),
		Name: s.Name,
		Spec: dbSyntheticSpecToV1(&s.Spec),
	}
}

func dbSyntheticSpecToV1(spec *db.SyntheticSpec) *apiv1.SyntheticSpec {
	var variables []*apiv1.Variable
	for _, v := range spec.Variables {
		variables = append(variables, &apiv1.Variable{
			Name:  v.Name,
			Value: v.Value,
		})
	}

	var steps []*apiv1.Step
	for _, s := range spec.Steps {
		steps = append(steps, dbStepToV1(s))
	}

	return &apiv1.SyntheticSpec{
		Variables: variables,
		Steps:     steps,
	}
}

func dbStepToV1(step *db.StepSpec) *apiv1.Step {
	return &apiv1.Step{
		Name: "TODO",
		Spec: &apiv1.StepSpec{
			Url:    step.URL,
			Method: http.MethodGet,
			Body:   *step.Body,
		},
	}
}
