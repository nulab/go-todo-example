package server

import (
	"github.com/nulab/go-todo-example/store"
	"encoding/json"
	"net/http"
)

var ds Store = &store.Datastore{}

// Store defines the datastore services
type Store interface {
	GetPendingTasks() []store.Task
	SaveTask(task store.Task) error
}

// GetPendingTasks returns pending tasks as a JSON response
func GetPendingTasks(w http.ResponseWriter, r *http.Request) {

	t := ds.GetPendingTasks()

	j, _ := json.Marshal(t)

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// AddTask handles requests for adding a new task.
// Return 201 if the task could be created
// Return 400 when JSON could not be decoded into a task or
// datastore returned an error or task title is empty
func AddTask(w http.ResponseWriter, r *http.Request) {
	var t store.Task

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if t.Title == "" {
		http.Error(w, "Title is missing", http.StatusBadRequest)
		return
	}

	if err := ds.SaveTask(t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}