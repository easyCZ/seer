package agent

// type SynteticSpec struct {
// 	Steps []StepSpec
// }

// type SyntheticResult struct {
// 	StepResults []*StepResult
// }

// func ExecuteSynthetic(ctx context.Context, spec *apiv1.SyntheticSpec) (*SyntheticResult, error) {
// 	vars := map[string]string{}

// 	var results []*StepResult
// 	for _, step := range spec.Steps {
// 		result, err := ExecuteStep(ctx, step, vars)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to execute step: %w", err)
// 		}

// 		results = append(results, result)

// 		// extend our variables with new extracts, we retain the previous ones as they are valid for all steps, unless overriden
// 		for key, value := range result.Extracts {
// 			vars[key] = value
// 		}
// 	}

// 	return &SyntheticResult{
// 		StepResults: results,
// 	}, nil
// }

// type ExtractSpec struct {
// 	Name       string
// 	Expression string
// }

// type StepSpec struct {
// 	URL      string
// 	Method   string
// 	Headers  http.Header
// 	Body     []byte
// 	Extracts []ExtractSpec
// }

// type StepResult struct {
// 	Body    []byte
// 	Headers http.Header
// 	Code    int

// 	Extracts map[string]string
// }

// func ExecuteStep(ctx context.Context, step *apiv1.Step, vars map[string]string) (*StepResult, error) {
// 	if vars == nil {
// 		vars = map[string]string{}
// 	}

// 	spec := step.Spec
// 	url := spec.Url
// 	for key, value := range vars {
// 		// try to replace the variable across most properties
// 		url = strings.ReplaceAll(url, fmt.Sprintf("{%s}", key), value)

// 		// TODO: replace other properties of the spec
// 	}

// 	req, err := http.NewRequestWithContext(ctx, spec.Method, url, bytes.NewBufferString(spec.Body))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to construct request: %w", err)
// 	}

// 	{
// 		d, err := httputil.DumpRequest(req, true)
// 		if err == nil {
// 			fmt.Println("Request")
// 			fmt.Println(string(d))
// 		}
// 	}

// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
// 	}

// 	{
// 		d, err := httputil.DumpResponse(resp, true)
// 		if err == nil {
// 			fmt.Println("Response")
// 			fmt.Println(string(d))
// 		}
// 	}

// 	return StepResultFromResponse(resp, spec)
// }

// func StepResultFromResponse(resp *http.Response, step *apiv1.StepSpec) (*StepResult, error) {
// 	if resp == nil {
// 		return nil, fmt.Errorf("nil response")
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("faield to read response body: %w", err)
// 	}

// 	var data interface{}
// 	if err := json.Unmarshal(body, &data); err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal body to JSON: %w", err)
// 	}

// 	extracts := map[string]string{}
// 	for _, e := range step.Extracts {
// 		query, err := gojq.Parse(e.Jql)
// 		if err != nil {
// 			fmt.Println("failed to parse extract expression", err)
// 			continue
// 		}

// 		iterator := query.Run(data)
// 		for {
// 			v, ok := iterator.Next()
// 			if !ok {
// 				break
// 			}

// 			if err, ok := v.(error); ok {
// 				fmt.Println("failed to query with jq", err)
// 				continue
// 			}
// 			extracts[e.Name] = fmt.Sprintf("%v", v)
// 		}

// 	}

// 	return &StepResult{
// 		Body:     body,
// 		Headers:  resp.Header.Clone(),
// 		Code:     resp.StatusCode,
// 		Extracts: extracts,
// 	}, nil
// }
