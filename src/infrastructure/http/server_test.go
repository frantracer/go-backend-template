package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	port = 3001
)

type unstructuredJSON = map[string]interface{}

func doRequest(reqMethod, reqURL string, reqBody unstructuredJSON) (respStatusCode int, respBody string, err error) {
	var marshalledBody []byte
	marshalledBody, err = json.Marshal(reqBody)
	if err != nil {
		return 0, "", err
	}

	req, err := http.NewRequest(reqMethod, reqURL, bytes.NewBuffer(marshalledBody))
	if err != nil {
		return 0, "", err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, "", err
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, "", err
	}

	err = resp.Body.Close()
	if err != nil {
		return 0, "", err
	}

	return resp.StatusCode, string(respBodyBytes), nil
}

func runServer() {
	ctx := context.Background()
	server := NewServer(ctx, NewHandler())

	readyChannel := make(chan struct{})
	go func() {
		err := server.ListenAndServe(port, readyChannel)
		if err != nil {
			panic("error while trying to run the server: " + err.Error())
		}
	}()

	<-readyChannel
}

func TestRun(t *testing.T) {
	runServer()
	baseURL := "http://localhost:" + strconv.Itoa(port)

	t.Run("Test general endpoints", func(t *testing.T) {
		t.Run("Health check returns an OK status", func(t *testing.T) {
			healthURL := baseURL + "/health"

			statusCode, _, err := doRequest(http.MethodGet, healthURL, nil)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, statusCode)
		})
	})
}
