package task

import "errors"

type Task struct {
	ID        int    `json:"id"`
	ListID    int    `json:"list_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// repository
type Repository interface {
	FindAll() []*Task
	FindByID(task_id int) (*Task, error)
	FindByListID(list_id int) []*Task
	Store(t *Task) error
}

var ErrUnknownTask = errors.New("unknown task")
