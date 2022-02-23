package agent

import (
	"context"
	"net/http"
	"testing"
	"time"

	apiv1 "github.com/easyCZ/seer/gen/v1"
	"github.com/stretchr/testify/require"
)

func TestRunner(t *testing.T) {
	runner := &Runner{
		ID:     "runner-id",
		Client: http.DefaultClient,
		debug:  true,
	}

	t.Run("single step success", func(t *testing.T) {
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

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		results, err := runner.Execute(ctx, synthetic)
		require.NoError(t, err)
		require.Len(t, results, 1)

		res := results[0]
		require.EqualValues(t, http.StatusOK, res.GetResponse().Status)

	})

	t.Run("single step failure", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		results, err := runner.Execute(ctx, &apiv1.Synthetic{
			Id:   "some-id",
			Name: "some-name",
			Spec: &apiv1.SyntheticSpec{
				Steps: []*apiv1.Step{
					{
						Name: "first",
						Spec: &apiv1.StepSpec{
							Url:    "<>../././not-a-valid-url",
							Method: "not a method",
						},
					},
				},
			},
		})
		require.NoError(t, err)
		require.Len(t, results, 1)

		require.EqualValues(t, "Failed to construct HTTP request", results[0].GetError().Message)
	})

}

const userJSON = `[
	{
		"id": 1,
		"name": "Leanne Graham",
		"username": "Bret",
		"email": "Sincere@april.biz",
		"phone": "1-770-736-8031 x56442",
		"website": "hildegard.org"
	}
]`

func TestEvaluteExtracts(t *testing.T) {
	resp := &apiv1.Response{
		Status: http.StatusOK,
		Headers: []*apiv1.Header{
			{Key: "Content-Type", Value: "application/json"},
			{Key: "X-Custom-Foo", Value: "henry has many sheep"},
		},
		Body: userJSON,
	}

	vars, err := evaluteExtracts([]*apiv1.Extract{
		{
			Name: "first",
			From: &apiv1.Extract_Header{
				Header: &apiv1.HeaderExtract{
					HeaderName: "Content-Type",
				},
			},
		},
	}, resp)
	require.NoError(t, err)
	require.Equal(t, []*apiv1.Variable{
		{Name: "first", Value: "application/json"},
	}, vars)
}
