package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSynthetic(t *testing.T) {
	u, err := url.Parse("https://jsonplaceholder.typicode.com/posts")
	require.NoError(t, err)

	headers := http.Header{}
	headers.Add("Accept", "application/json")

	spec := StepSpec{
		URL:     u,
		Method:  http.MethodGet,
		Headers: headers,
		Extracts: []ExtractSpec{
			{Name: "TEST", Expression: ".[0].userId"},
		},
	}

	result, err := Execute(context.Background(), spec)
	require.NoError(t, err)

	fmt.Println(string(result.Body))
	fmt.Println(result.Extracts)
	require.Equal(t, map[string]string{
		"TEST": "1",
	}, result.Extracts)
}
