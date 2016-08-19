package server

import (
	"github.com/nulab/go-todo-example/store"
	"encoding/json"
	"net/http"
)

var ds = &store.Datastore{}

// GetPendingTasks returns pending tasks as a JSON response
func GetPendingTasks(w http.ResponseWriter, r *http.Request) {

	t := ds.GetPendingTasks()

	j, _ := json.Marshal(t)

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
