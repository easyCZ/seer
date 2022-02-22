package agent

import (
	"context"
	"net/http"
	"testing"
	"time"

	apiv1 "github.com/easyCZ/qfy/gen/v1"
	"github.com/stretchr/testify/require"
)

func TestRunner(t *testing.T) {
	synthetic := &apiv1.Synthetic{
		Id:   "123",
		Name: "synthetic-name",
		Spec: &apiv1.SyntheticSpec{
			Variables: []*apiv1.Variable{
				{Name: "API_ROOT", Value: "https://jsonplaceholder.typicode.com"},
			},
			Steps: []*apiv1.Step{
				{
					Name: "List posts",
					Spec: &apiv1.StepSpec{
						Url:    "{API_ROOT}/posts",
						Method: http.MethodGet,
					},
				},
			},
		},
	}

	runner := &Runner{
		ID:     "runner-id",
		Client: http.DefaultClient,
		debug:  true,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	results, err := runner.Execute(ctx, synthetic)
	require.NoError(t, err)
	require.Len(t, results, 1)

	res := results[0]
	switch op := res.Outcome.(type) {
	case *apiv1.StepResult_Response:
		require.EqualValues(t, http.StatusOK, op.Response.Status)
	default:
		require.Fail(t, "expected succesful response")
	}

}
