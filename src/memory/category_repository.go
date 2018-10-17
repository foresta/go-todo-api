package memory

import (
	"sync"

	"github.com/foresta/go-todo-api/src/category"
)

type categoryRepository struct {
	mtx        sync.RWMutex
	categories map[int]*category.Category
}

func NewCategoryRepository() category.Repository {
	return &categoryRepository{
		categories: make(map[int]*category.Category),
	}
}

func (repo *categoryRepository) FindAll() []*category.Category {
	repo.mtx.RLock()
	defer repo.mtx.RUnlock()

	c := make([]*category.Category, 0, len(repo.categories))
	for _, val := range repo.categories {
		c = append(c, val)
	}
	return c
}

func (repo *categoryRepository) FindByID(id int) (*category.Category, error) {
	repo.mtx.RLock()
	defer repo.mtx.RUnlock()

	val, ok := repo.categories[id]
	if ok {
		return val, nil
	}

	return nil, category.ErrUnknownCategory
}

func (repo *categoryRepository) Store(c *category.Category) error {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()
	repo.categories[c.ID] = c
	return nil
}
