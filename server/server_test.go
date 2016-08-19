package server

import (
	"bytes"
	"errors"
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
	name     string
	saveFunc func(task store.Task) error
	body     []byte
	want     int
}{
	{
		name: "should add new task from JSON",
		body: []byte(`{"Title":"Buy bread for breakfast."}`),
		want: http.StatusCreated,
	},
	{
		name: "should response bad argument when JSON could not be handled",
		body: []byte(""),
		want: http.StatusBadRequest,
	},
	{
		name: "should response bad argument when datastore returns an error",
		saveFunc: func(task store.Task) error {
			return errors.New("datastore error")
		},
		body: []byte(`{"Title":"Buy bread for breakfast."}`),
		want: http.StatusBadRequest,
	},
	{
		name: "should response bad argument when task title is emtpy",
		body: []byte(`{"Title":""}`),
		want: http.StatusBadRequest,
	},
}

func TestAddTask(t *testing.T) {

	t.Log("AddTask")

	for _, testcase := range addTaskTests {

		t.Log(testcase.name)

		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(testcase.body))

		defer func() { ds = &store.Datastore{} }()

		ds = &mockedStore{
			SaveTaskFunc: testcase.saveFunc,
		}

		AddTask(rec, req)

		if rec.Code != testcase.want {
			t.Errorf("KO => Got %d wanted %d", rec.Code, testcase.want)
		}
	}
}
type mockedStore struct {
	SaveTaskFunc func(task store.Task) error
}

func (ms *mockedStore) GetPendingTasks() []store.Task {
	return []store.Task{
		{1, "Do housework", false},
		{2, "Buy milk", false},
	}
}

func (ms *mockedStore) SaveTask(task store.Task) error {
	if ms.SaveTaskFunc != nil {
		return ms.SaveTaskFunc(task)
	}
	return nil
}
