package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/itchyny/gojq"
)

type ExtractSpec struct {
	Name       string
	Expression string
}

type StepSpec struct {
	URL      *url.URL
	Method   string
	Headers  http.Header
	Body     []byte
	Extracts []ExtractSpec
}

type StepResult struct {
	Body    []byte
	Headers http.Header
	Code    int

	Extracts map[string]string
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

	return StepResultFromResponse(resp, step)
}

func StepResultFromResponse(resp *http.Response, step StepSpec) (*StepResult, error) {
	if resp == nil {
		return nil, fmt.Errorf("nil response")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("faield to read response body: %w", err)
	}

	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal body to JSON: %w", err)
	}

	extracts := map[string]string{}
	for _, e := range step.Extracts {
		query, err := gojq.Parse(e.Expression)
		if err != nil {
			fmt.Println("failed to parse extract expression", err)
			continue
		}

		iterator := query.Run(data)
		for {
			v, ok := iterator.Next()
			if !ok {
				break
			}

			if err, ok := v.(error); ok {
				fmt.Println("failed to query with jq", err)
				continue
			}
			extracts[e.Name] = fmt.Sprintf("%v", v)
		}

	}

	return &StepResult{
		Body:     body,
		Headers:  resp.Header.Clone(),
		Code:     resp.StatusCode,
		Extracts: extracts,
	}, nil
}
