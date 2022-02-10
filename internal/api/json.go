package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, r *http.Request, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	enc := json.NewEncoder(w)
	if err := enc.Encode(data); err != nil {
		return fmt.Errorf("failed to write json to response: %w", err)
	}

	return nil
}
