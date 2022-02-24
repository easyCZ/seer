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
	"regexp"
	"strings"

	apiv1 "github.com/easyCZ/seer/gen/v1"

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
				Key:   strings.ToLower(name),
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

func evaluteExtracts(extracts []*apiv1.Extract, resp *apiv1.Response) ([]*apiv1.Variable, error) {
	var vars []*apiv1.Variable

	for _, e := range extracts {
		switch t := e.GetFrom().(type) {
		case *apiv1.Extract_Body:
			logger.Printf("body %v", t)

			return nil, nil
		case *apiv1.Extract_Header:
			logger.Printf("header %v", t)

			for _, header := range resp.Headers {
				if header.Key == t.Header.HeaderName {
					val, matched, err := extractFromString(header.Value, t.Header.Query)
					if err != nil {
						return nil, fmt.Errorf("failed to extract from header %s: %w", header.Key, err)
					}
					if matched {
						vars = append(vars, &apiv1.Variable{
							Name:  e.Name,
							Value: val,
						})
					}
				}
			}
		default:
			return nil, fmt.Errorf("invalid extract from definition")
		}
	}

	return vars, nil
}

func extractFromString(s string, q *apiv1.ExtractQuery) (string, bool, error) {
	fmt.Printf("extract query %v", q)
	if q == nil {
		return s, true, nil
	}

	switch t := q.Expression.(type) {
	case *apiv1.ExtractQuery_Jq:
		logger.Printf("JQL %v", t)
		return "", false, nil

	case *apiv1.ExtractQuery_Regexp:
		exp, err := regexp.Compile(t.Regexp)
		if err != nil {
			return "", false, fmt.Errorf("failed to compile regular expression: %w", err)
		}

		matches := exp.FindStringSubmatch(s)
		logger.Printf("Regexp matches: %v", matches)
		if len(matches) > 1 {
			return matches[1], true, nil
		}
		return "", false, nil

	default:
		return s, true, nil
	}
}
