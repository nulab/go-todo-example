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

func TestSaveNewTask(t *testing.T) {
	t.Log("SaveTask")

	ds := Datastore{}

	task := Task{Title: "Buy milk"}

	want := []Task{
		{1, "Buy milk", false},
	}

	t.Log("should save the new task in the store")

	ds.SaveTask(task)

	if !reflect.DeepEqual(ds.tasks, want) {
		t.Errorf("=> Got %v wanted %v", ds.tasks, want)
	}
}