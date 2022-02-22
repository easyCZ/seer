package agent

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	apiv1 "github.com/easyCZ/qfy/gen/v1"

	"sync/atomic"
)

type Runner struct {
	ID string

	Client *http.Client

	isExcecuting int32

	// When enabled, http requests are debug logged to stdout
	debug bool
}

var (
	logger = log.New(os.Stderr, "runner", log.Default().Flags())
)

func (r *Runner) Execute(ctx context.Context, syn *apiv1.Synthetic) ([]*apiv1.StepResult, error) {
	if !atomic.CompareAndSwapInt32(&r.isExcecuting, 0, 1) {
		return nil, fmt.Errorf("running is already executing")
	}

	defer (func() {
		atomic.CompareAndSwapInt32(&r.isExcecuting, 1, 0)
	})()

	spec := syn.Spec
	vars := spec.Variables

	var results []*apiv1.StepResult
	for i, step := range spec.Steps {
		logger.Printf("Executing step #%d - %s", i, step.Name)
		result, err := r.executeStep(ctx, vars, step)
		if err != nil {
			// TODO: Will need to package the errors into the Result so it can be observed.
			return nil, fmt.Errorf("failed to execute step: %w", err)
		}

		results = append(results, result)
		vars = result.Variables
	}

	return results, nil
}

func (r *Runner) executeStep(ctx context.Context, vars []*apiv1.Variable, step *apiv1.Step) (*apiv1.StepResult, error) {
	spec := step.Spec

	url := replaceStringWithVars(spec.Url, vars)
	body := replaceStringWithVars(spec.Body, vars)

	request, err := http.NewRequestWithContext(ctx, spec.Method, url, bytes.NewBufferString(body))
	if err != nil {
		return &apiv1.StepResult{
			StepName: step.Name,
			Outcome: &apiv1.StepResult_Error{
				Error: &apiv1.StepError{
					Message: "Failed to construct HTTP request",
					Details: err.Error(),
				},
			},
		}, nil
	}

	if r.debug {
		d, err := httputil.DumpRequest(request, true)
		if err == nil {
			fmt.Println("Request")
			fmt.Println(string(d))
		}
	}

	httpResp, err := r.Client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}

	if r.debug {
		d, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println("Response")
			fmt.Println(string(d))
		}
	}

	response, err := convertResponse(httpResp)
	if err != nil {
		return nil, fmt.Errorf("failed to convert response: %w", err)
	}

	return &apiv1.StepResult{
		StepName: step.Name,
		Outcome: &apiv1.StepResult_Response{
			Response: response,
		},
	}, nil
}

func replaceStringWithVars(s string, vars []*apiv1.Variable) string {
	for _, v := range vars {
		s = strings.ReplaceAll(s, fmt.Sprintf("{%s}", v.Name), v.Value)
	}

	return s
}

func convertResponse(r *http.Response) (*apiv1.Response, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if err := r.Body.Close(); err != nil {
		return nil, fmt.Errorf("failed to close response body: %w", err)
	}

	var headers []*apiv1.Header
	for name := range r.Header {
		for _, value := range r.Header.Values(name) {
			headers = append(headers, &apiv1.Header{
				Key:   name,
				Value: value,
			})
		}
	}

	return &apiv1.Response{
		Status:  int32(r.StatusCode),
		Body:    string(body),
		Headers: headers,
	}, nil
}
