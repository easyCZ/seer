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
	u, err := url.Parse("https://reddit.com/.json")
	require.NoError(t, err)

	headers := http.Header{}
	headers.Add("Accept", "application/json")

	spec := StepSpec{
		URL:     u,
		Method:  http.MethodGet,
		Headers: headers,
	}

	result, err := Execute(context.Background(), spec)
	require.NoError(t, err)

	fmt.Println(result)
	fmt.Println(string(result.Body))

}
