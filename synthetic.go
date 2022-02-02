package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type StepSpec struct {
	URL     *url.URL
	Method  string
	Headers http.Header
	Body    []byte
}

type StepResult struct {
	Body    []byte
	Headers http.Header
	Code    int
}

func Execute(ctx context.Context, step StepSpec) (*StepResult, error) {

	req, err := http.NewRequestWithContext(ctx, step.Method, step.URL.String(), bytes.NewBuffer(step.Body))
	if err != nil {
		return nil, fmt.Errorf("failed to construct request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}

	return StepResultFromResponse(resp)
}

func StepResultFromResponse(resp *http.Response) (*StepResult, error) {
	if resp == nil {
		return nil, fmt.Errorf("nil response")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("faield to read response body: %w", err)
	}

	return &StepResult{
		Body:    body,
		Headers: resp.Header.Clone(),
		Code:    resp.StatusCode,
	}, nil
}
