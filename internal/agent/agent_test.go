package agent

import "testing"

func TestAgent(t *testing.T) {
	a := &Agent{URL: "localhost:3000"}

	a.Run()
}
