package functions

import (
	"encoding/json"
	"net/http"

	"github.com/frantacer/go-backend-template/src/application/handlers"
	"github.com/frantacer/go-backend-template/src/domain"
)

type InsertTaskBody struct {
	Message string `json:"message"`
}

type TaskJSON struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

func InsertTaskHTTPHandler(appHandler handlers.InsertTaskHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body InsertTaskBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		task, _ := appHandler.Handle(r.Context(), handlers.InsertTaskCommand{Message: body.Message})
		data, _ := json.Marshal(mapTaskFromDomainToJSON(task))

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(data)
	}
}

func FindTasksHTTPHandler(appHandler handlers.FindTasksHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, _ := appHandler.Handle(r.Context())

		items := make([]TaskJSON, len(tasks))
		for i := range tasks {
			items[i] = mapTaskFromDomainToJSON(tasks[i])
		}
		data, _ := json.Marshal(items)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(data)
	}
}

func mapTaskFromDomainToJSON(task domain.Task) TaskJSON {
	return TaskJSON{
		ID:      task.ID().String(),
		Message: task.Message(),
	}
}
