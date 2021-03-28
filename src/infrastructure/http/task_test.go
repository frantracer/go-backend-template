package http_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	appHandlers "github.com/frantacer/go-backend-template/src/application/handlers"
	"github.com/frantacer/go-backend-template/src/domain"
	httpx "github.com/frantacer/go-backend-template/src/infrastructure/http"
)

func TestInsertTask(t *testing.T) {
	t.Run("Given the insert task handler does not fail,"+
		"when the create task endpoint is called,"+
		"then the new task is returned as JSON",
		func(t *testing.T) {
			task := domain.NewTask(uuid.New(), "my task")
			appHandler := InsertTaskHandlerMock{
				HandleFunc: func(ctx context.Context, cmd appHandlers.InsertTaskCommand) (domain.Task, error) {
					return task, nil
				},
			}
			httpHandler := httpx.NewHandler(httpx.ApplicationHandlers{
				InsertTaskHandler: &appHandler,
			})
			body, _ := json.Marshal(httpx.InsertTaskRequestBody{Message: task.Message()})
			req := httptest.NewRequest("POST", "/tasks", bytes.NewReader(body))
			expectedResponse := httpx.InsertTaskResponseBody{
				ID:      task.ID().String(),
				Message: task.Message(),
			}

			resp := sendRequest(t, httpHandler, req)

			require.Equal(t, http.StatusOK, resp.StatusCode)

			data := mustParseInsertTaskResponse(resp.Body)
			require.Equal(t, expectedResponse, data)
		})
}

func TestRetrieveTasks(t *testing.T) {
	t.Run("Given the handler returns two tasks,"+
		"when the list of tasks endpoint is called,"+
		"then it returns the two tasks as a JSON array",
		func(t *testing.T) {
			task := domain.NewTask(uuid.New(), "my task")
			appHandler := FindTasksHandlerMock{
				HandleFunc: func(ctx context.Context) ([]domain.Task, error) {
					return []domain.Task{task}, nil
				},
			}
			httpHandler := httpx.NewHandler(httpx.ApplicationHandlers{
				FindTasksHandler: &appHandler,
			})
			req := httptest.NewRequest("GET", "/tasks", nil)
			expectedResponse := httpx.FindTasksResponseBody{
				{
					ID:      task.ID().String(),
					Message: task.Message(),
				},
			}

			resp := sendRequest(t, httpHandler, req)

			require.Equal(t, http.StatusOK, resp.StatusCode)
			fmt.Println(resp.Body)

			data := mustParseFindTasksResponse(resp.Body)
			require.Equal(t, expectedResponse, data)
		})
}

func mustParseInsertTaskResponse(body httpx.ResponseBody) httpx.InsertTaskResponseBody {
	response := httpx.InsertTaskResponseBody{}
	_ = json.Unmarshal(body.Data, &response)
	return response
}

func mustParseFindTasksResponse(body httpx.ResponseBody) httpx.FindTasksResponseBody {
	response := httpx.FindTasksResponseBody{}
	_ = json.Unmarshal(body.Data, &response)
	return response
}
