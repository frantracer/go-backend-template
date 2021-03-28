package http

import (
	"encoding/json"
	"net/http"
)

type HealthCheckResponse struct {
	Status string `json:"status"`
}

func HealthCheckHTTPFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := HealthCheckResponse{Status: "running"}
		data, err := json.Marshal(response)
		if err != nil {
			return
		}

		sendResponse(w, data)
	}
}
