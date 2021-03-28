package http

import (
	"encoding/json"
	"net/http"
)

type ResponseBody struct {
	Data  json.RawMessage `json:"data,omitempty"`
	Error string          `json:"error,omitempty"`
}

func sendResponse(w http.ResponseWriter, data json.RawMessage) {
	response := ResponseBody{
		Data: data,
	}

	body, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(body)
	w.WriteHeader(http.StatusOK)
}

func sendError(w http.ResponseWriter, errorMsg string, errorCode int) {
	response := ResponseBody{
		Error: errorMsg,
	}

	body, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(body)
	w.WriteHeader(errorCode)
}
