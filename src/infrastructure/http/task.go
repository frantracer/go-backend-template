package http

import (
	"encoding/json"
	"net/http"

	"github.com/frantacer/go-backend-template/src/application/handlers"
	"github.com/frantacer/go-backend-template/src/domain"
)

type TaskJSON struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type InsertTaskRequestBody struct {
	Message string `json:"message"`
}

type InsertTaskResponseBody = TaskJSON

func InsertTaskHTTPFunc(appHandler handlers.InsertTaskHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body InsertTaskRequestBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendError(w, err.Error(), http.StatusBadRequest)
			return
		}

		task, _ := appHandler.Handle(r.Context(), handlers.InsertTaskCommand{Message: body.Message})
		data, _ := json.Marshal(mapTaskFromDomainToJSON(task))

		sendResponse(w, data)
	}
}

type FindTasksResponseBody = []TaskJSON

func FindTasksHTTPFunc(appHandler handlers.FindTasksHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, _ := appHandler.Handle(r.Context())

		items := make(FindTasksResponseBody, len(tasks))
		for i := range tasks {
			items[i] = mapTaskFromDomainToJSON(tasks[i])
		}
		data, _ := json.Marshal(items)

		sendResponse(w, data)
	}
}

func mapTaskFromDomainToJSON(task domain.Task) TaskJSON {
	return TaskJSON{
		ID:      task.ID().String(),
		Message: task.Message(),
	}
}
