package store

import (
	"reflect"
	"testing"
)

func TestGetPendingTasks(t *testing.T) {
	t.Log("GetPendingTasks")

	ds := Datastore{
		tasks: []Task{
			{1, "Do housework", true},
			{2, "Buy milk", false},
		},
	}

	want := []Task{ds.tasks[1]}

	t.Log("should return the tasks which need to be completed")

	if got := ds.GetPendingTasks(); !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v wanted %v", got, want)
	}
}

var saveTaskTests = []struct {
	name string
	ds   *Datastore
	task Task
	want []Task
}{
	{
		name: "should save the new task in the datastore",
		ds:   &Datastore{},
		task: Task{Title: "Buy milk"},
		want: []Task{
			{1, "Buy milk", false},
		},
	},
	{
		name: "should update the existing task in the store",
		ds: &Datastore{
			tasks: []Task{
				{1, "Buy milk", false},
			},
		},
		task: Task{1, "Buy milk", true},
		want: []Task{
			{1, "Buy milk", true},
		},
	},
}

func TestSaveTask(t *testing.T) {
	t.Log("SaveTask")

	for _, testcase := range saveTaskTests {
		t.Log(testcase.name)
		testcase.ds.SaveTask(testcase.task)

		if !reflect.DeepEqual(testcase.ds.tasks, testcase.want) {
			t.Errorf("=> Got %v wanted %v", testcase.ds.tasks, testcase.want)
		}
	}
}