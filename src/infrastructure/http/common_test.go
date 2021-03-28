package http_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	httpx "github.com/frantacer/go-backend-template/src/infrastructure/http"
)

type CustomResponseHTTP struct {
	StatusCode int
	Body       httpx.ResponseBody
}

func sendRequest(t *testing.T, handler http.Handler, request *http.Request) CustomResponseHTTP {
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)
	responseHTTP := recorder.Result()

	bodyBytes, err := ioutil.ReadAll(responseHTTP.Body)
	require.NoError(t, err)
	_ = responseHTTP.Body.Close()

	response := CustomResponseHTTP{}
	response.StatusCode = responseHTTP.StatusCode
	_ = json.Unmarshal(bodyBytes, &response.Body)

	return response
}
