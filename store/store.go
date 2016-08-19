package store

import "errors"

// Task job to be done or completed
type Task struct {
	ID    int    `json:"id"`    // identifier of the task
	Title string `json:"title"` // Title of the task
	Done  bool   `json:"done"`  // If task is completed or not
}

// Datastore manages a list of tasks stored in memory
type Datastore struct {
	tasks  []Task
	lastID int // lastID is incremented for each new stored task
}

// GetPendingTasks returns all the tasks which need to be done
func (ds *Datastore) GetPendingTasks() []Task {
	var pendingTasks []Task
	for _, task := range ds.tasks {
		if !task.Done {
			pendingTasks = append(pendingTasks, task)
		}
	}
	return pendingTasks
}

// ErrTaskNotFound is returned when a Task ID is not found
var ErrTaskNotFound = errors.New("Task was not found")

// SaveTask should store the task in the datastore if the task
// is new else update it. A Task Not Found error is returned
// when the task ID does not exist
func (ds *Datastore) SaveTask(task Task) error {
	if task.ID == 0 {
		ds.lastID++
		task.ID = ds.lastID
		ds.tasks = append(ds.tasks, task)
		return nil
	}

	for i, t := range ds.tasks {
		if t.ID == task.ID {
			ds.tasks[i] = task
			return nil
		}
	}

	return ErrTaskNotFound
}