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