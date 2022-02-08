package main

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStep(t *testing.T) {
	headers := http.Header{}
	headers.Add("Accept", "application/json")

	spec := StepSpec{
		URL:     "https://jsonplaceholder.typicode.com/posts",
		Method:  http.MethodGet,
		Headers: headers,
		Extracts: []ExtractSpec{
			{Name: "TEST", Expression: ".[0].userId"},
		},
	}

	result, err := ExecuteStep(context.Background(), spec, nil)
	require.NoError(t, err)

	fmt.Println(string(result.Body))
	fmt.Println(result.Extracts)
	require.Equal(t, map[string]string{
		"TEST": "1",
	}, result.Extracts)
}

func TestSynthetic(t *testing.T) {
	spec := SynteticSpec{
		Steps: []StepSpec{
			{
				URL:     "https://jsonplaceholder.typicode.com/posts",
				Method:  http.MethodGet,
				Headers: headers(t, "Accept", "application/json"),
				Extracts: []ExtractSpec{
					{Name: "POST_ID", Expression: ".[0].id"}, // extract the post ID
				},
			},
			{
				URL:     "https://jsonplaceholder.typicode.com/posts/{POST_ID}/comments",
				Method:  http.MethodGet,
				Headers: headers(t, "Accept", "application/json"),
				Extracts: []ExtractSpec{
					{Name: "COMMENT_EMAIL", Expression: ".[0].email"}, // extract the comment email
				},
			},
		},
	}

	results, err := ExecuteSynthetic(context.Background(), spec)
	require.NoError(t, err)

	require.EqualValues(t, map[string]string{
		"COMMENT_EMAIL": "Eliseo@gardner.biz",
	}, results.StepResults[len(results.StepResults)-1].Extracts)
}

func headers(t *testing.T, args ...string) http.Header {
	t.Helper()

	if len(args)%2 != 0 {
		require.Fail(t, "header arguments are not of even length")
	}

	h := http.Header{}
	for i := 0; i < len(args); i += 2 {
		key := args[i]
		value := args[i+1]
		h.Add(key, value)
	}

	return h
}
