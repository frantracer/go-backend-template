package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	httpx "github.com/frantacer/go-backend-template/src/infrastructure/http"
)

func TestHealthHandler(t *testing.T) {
	t.Run("Health check returns an OK status", func(t *testing.T) {
		handle := httpx.NewHandler(httpx.ApplicationHandlers{})
		req := httptest.NewRequest("GET", "/health", nil)
		resp := sendRequest(t, handle, req)

		require.Equal(t, http.StatusOK, resp.StatusCode)

		data := mustParseHealthCheckResponse(t, resp.Body)
		require.Equal(t, "running", data.Status)
	})
}

func mustParseHealthCheckResponse(t *testing.T, body httpx.ResponseBody) httpx.HealthCheckResponse {
	response := httpx.HealthCheckResponse{}
	err := json.Unmarshal(body.Data, &response)
	require.NoError(t, err)
	return response
}
