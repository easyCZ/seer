package agent

// func TestStep(t *testing.T) {
// 	headers := http.Header{}
// 	headers.Add("Accept", "application/json")

// 	spec := &apiv1.StepSpec{
// 		Url:    "https://jsonplaceholder.typicode.com/posts",
// 		Method: http.MethodGet,
// 		// Headers: headers,
// 		Extracts: []*apiv1.Extract{
// 			{Name: "TEST", Jql: ".[0].userId"},
// 		},
// 	}

// 	result, err := ExecuteStep(context.Background(), &apiv1.Step{
// 		Name: "Test",
// 		Spec: spec,
// 	}, nil)
// 	require.NoError(t, err)

// 	fmt.Println(string(result.Body))
// 	fmt.Println(result.Extracts)
// 	require.Equal(t, map[string]string{
// 		"TEST": "1",
// 	}, result.Extracts)
// }

// func TestSynthetic(t *testing.T) {
// 	spec := &apiv1.SyntheticSpec{
// 		Steps: []*apiv1.Step{
// 			{
// 				Name: "Fetch posts",
// 				Spec: &apiv1.StepSpec{
// 					Url:    "https://jsonplaceholder.typicode.com/posts",
// 					Method: http.MethodGet,
// 					// Headers: headers(t, "Accept", "application/json"),
// 					Extracts: []*apiv1.Extract{
// 						{Name: "POST_ID", Jql: ".[0].id"}, // extract the post ID
// 					},
// 				},
// 			},
// 			{
// 				Name: "Get comments for first post",
// 				Spec: &apiv1.StepSpec{
// 					Url:    "https://jsonplaceholder.typicode.com/posts/{POST_ID}/comments",
// 					Method: http.MethodGet,
// 					// Headers: headers(t, "Accept", "application/json"),
// 					Extracts: []*apiv1.Extract{
// 						{Name: "COMMENT_EMAIL", Jql: ".[0].email"}, // extract the comment email
// 					},
// 				},
// 			},
// 		},
// 	}

// 	results, err := ExecuteSynthetic(context.Background(), spec)
// 	require.NoError(t, err)

// 	require.EqualValues(t, map[string]string{
// 		"COMMENT_EMAIL": "Eliseo@gardner.biz",
// 	}, results.StepResults[len(results.StepResults)-1].Extracts)
// }
