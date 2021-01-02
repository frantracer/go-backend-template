package functions

import (
	"encoding/json"
	"net/http"
)

type healthCheckResponse struct {
	Status string `json:"status"`
}

// HealthCheckHandler is the handler for health endpoint.
func HealthCheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := healthCheckResponse{Status: "running"}
		data, err := json.Marshal(response)
		if err != nil {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(data)
	}
}
