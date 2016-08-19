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

var addTaskTests = []struct {
	name string
	body []byte
	want int
}{
	{
		name: "should add new task from JSON",
		body: []byte(`{"Title":"Buy bread for breakfast."}`),
		want: http.StatusCreated,
	},
}

func TestAddTask(t *testing.T) {

	t.Log("AddTask")

	for _, testcase := range addTaskTests {

		t.Log(testcase.name)

		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(testcase.body))

		defer func() { ds = &store.Datastore{} }()

		ds = &mockedStore{}

		AddTask(rec, req)

		if rec.Code != testcase.want {
			t.Errorf("KO => Got %d wanted %d", rec.Code, testcase.want)
		}
	}
}

type mockedStore struct{}

func (ms *mockedStore) GetPendingTasks() []store.Task {
	return []store.Task{
		{1, "Do housework", false},
		{2, "Buy milk", false},
	}
}
