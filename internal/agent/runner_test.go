package agent

import (
	"context"
	"fmt"
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

		ctx, cancel := context.WithTimeout(context.Background(), 95*time.Second)
		defer cancel()

		results, err := runner.Execute(ctx, synthetic)
		require.NoError(t, err)
		require.Len(t, results, 1)

		res := results[0]
		require.EqualValues(t, http.StatusOK, res.GetResponse().Status)
	})

	t.Run("multi-step success", func(t *testing.T) {
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
							Extracts: []*apiv1.Extract{
								{
									Name: "FIRST_POST_ID",
									From: &apiv1.Extract_Body{
										Body: &apiv1.BodyExtract{
											Query: &apiv1.ExtractQuery{
												Expression: &apiv1.ExtractQuery_Jq{
													Jq: ".[0].id",
												},
											},
										},
									},
								},
							},
						},
					},
					{
						Name: "Get first comment for first post",
						Spec: &apiv1.StepSpec{
							Url:    "{API_ROOT}/posts/{FIRST_POST_ID}/comments",
							Method: http.MethodGet,
							Extracts: []*apiv1.Extract{
								{
									Name: "FIRST_POST_FIRST_COMMENT",
									From: &apiv1.Extract_Body{
										Body: &apiv1.BodyExtract{
											Query: &apiv1.ExtractQuery{
												Expression: &apiv1.ExtractQuery_Jq{
													Jq: ".[0].body",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		results, err := runner.Execute(ctx, synthetic)
		require.NoError(t, err)
		require.Len(t, results, 2)

		res := results[1]
		fmt.Println("variables", res.Variables)
		v, found := findVariableByKey(t, res.Variables, "FIRST_POST_FIRST_COMMENT")
		require.True(t, found)

		require.EqualValues(t, "laudantium enim quasi est quidem magnam voluptate ipsam eos\ntempora quo necessitatibus\ndolor quam autem quasi\nreiciendis et nam sapiente accusantium", v.Value)

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

	t.Run("extract whole header when query not specified", func(t *testing.T) {
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
	})

	t.Run("extract a specific part of the header when query specified", func(t *testing.T) {
		vars, err := evaluteExtracts([]*apiv1.Extract{
			{
				Name: "first",
				From: &apiv1.Extract_Header{
					Header: &apiv1.HeaderExtract{
						HeaderName: "Content-Type",
						Query: &apiv1.ExtractQuery{
							Expression: &apiv1.ExtractQuery_Regexp{
								Regexp: "application/(.*)",
							},
						},
					},
				},
			},
		}, resp)
		require.NoError(t, err)
		require.Equal(t, []*apiv1.Variable{
			{Name: "first", Value: "json"},
		}, vars)
	})

	t.Run("extract using JQ from body", func(t *testing.T) {
		vars, err := evaluteExtracts([]*apiv1.Extract{
			{
				Name: "first",
				From: &apiv1.Extract_Body{
					Body: &apiv1.BodyExtract{
						Query: &apiv1.ExtractQuery{
							Expression: &apiv1.ExtractQuery_Jq{
								Jq: ".[0].id",
							},
						},
					},
				},
			},
		}, resp)

		require.NoError(t, err)
		require.Equal(t, []*apiv1.Variable{
			{Name: "first", Value: "1"},
		}, vars)
	})
}

func findVariableByKey(t *testing.T, vars []*apiv1.Variable, key string) (*apiv1.Variable, bool) {
	for _, v := range vars {
		if v.Name == key {
			return v, true
		}
	}

	return nil, false
}
