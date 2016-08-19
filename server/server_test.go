package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nulab/go-todo-example/store"
	"bytes"
)

func TestGetPendingTasks(t *testing.T) {

	t.Log("GetPendingTasks")

	t.Log("should return pending tasks as JSON")

	rec := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/tasks/pending", nil)

	// The datastore is restored at the end of the test
	defer func() { ds = &store.Datastore{} }()

	ds = &mockedStore{}

	GetPendingTasks(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("KO => Got %d wanted %d", rec.Code, http.StatusOK)
	}

	want := "[{\"id\":1,\"title\":\"Do housework\",\"done\":false},{\"id\":2,\"title\":\"Buy milk\",\"done\":false}]"
	if got := rec.Body.String(); got != want {
		t.Errorf("KO => Got %s wanted %s", got, want)
	}
}

func TestAddTask(t *testing.T) {

	t.Log("AddTask")

	t.Log("should add new task from JSON")

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer([]byte(`{"Title":"Buy bread for breakfast."}`)))

	defer func() { ds = &store.Datastore{} }()

	ds = &mockedStore{}

	AddTask(rec, req)

	wantCode := http.StatusCreated
	if rec.Code != wantCode {
		t.Errorf("KO => Got %d wanted %d", rec.Code, wantCode)
	}
}

type mockedStore struct{}

func (ms *mockedStore) GetPendingTasks() []store.Task {
	return []store.Task{
		{1, "Do housework", false},
		{2, "Buy milk", false},
	}
}
