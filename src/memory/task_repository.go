package memory

import (
	"sync"

	"github.com/foresta/go-todo-api/src/task"
)

type taskRepository struct {
	mtx   sync.RWMutex
	tasks map[int]*task.Task
}

func NewTaskRepository() task.Repository {
	return &taskRepository{
		tasks: make(map[int]*task.Task),
	}
}

func (repo *taskRepository) FindAll() []*task.Task {
	repo.mtx.RLock()
	defer repo.mtx.RUnlock()

	t := make([]*task.Task, 0, len(repo.tasks))
	for _, val := range repo.tasks {
		t = append(t, val)
	}

	return t
}

func (repo *taskRepository) FindByID(id int) (*task.Task, error) {
	repo.mtx.RLock()
	defer repo.mtx.RUnlock()

	val, ok := repo.tasks[id]
	if ok {
		return val, nil
	}

	return nil, task.ErrUnknownTask
}

func (repo *taskRepository) FindByCategoryID(category_id int) []*task.Task {
	repo.mtx.RLock()
	defer repo.mtx.RUnlock()

	t := make([]*task.Task, 0)
	for _, val := range repo.tasks {
		if val.CategoryID == category_id {
			t = append(t, val)
		}
	}
	return t
}

func (repo *taskRepository) Store(t *task.Task) error {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()
	repo.tasks[t.ID] = t
	return nil
}
