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
}

// GetPendingTasks returns pending tasks as a JSON response
func GetPendingTasks(w http.ResponseWriter, r *http.Request) {

	t := ds.GetPendingTasks()

	j, _ := json.Marshal(t)

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// AddTask handles POST requests on /tasks.
// Return 201 if the task could be created
func AddTask(w http.ResponseWriter, r *http.Request) {
	var t store.Task

	json.NewDecoder(r.Body).Decode(&t)

	w.WriteHeader(http.StatusCreated)
}