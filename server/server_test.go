package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nulab/go-todo-example/store"
)

func TestGetPendingTasks(t *testing.T) {

	t.Log("GetPendingTasks")

	t.Log("should return pending tasks as JSON")

	rec := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/tasks/pending", nil)

	ds = &store.Datastore{
		Tasks: []store.Task{
			{1, "Do housework", false},
			{2, "Buy milk", false},
		},
	}

	GetPendingTasks(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("KO => Got %d wanted %d", rec.Code, http.StatusOK)
	}

	want := "[{\"id\":1,\"title\":\"Do housework\",\"done\":false},{\"id\":2,\"title\":\"Buy milk\",\"done\":false}]"
	if got := rec.Body.String(); got != want {
		t.Errorf("KO => Got %s wanted %s", got, want)
	}
}
