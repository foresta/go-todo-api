package category

import "errors"

type Category struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// repository
type Repository interface {
	FindAll() []*Category
	FindByID(id int) (*Category, error)
	Store(c *Category) error
}

// error
var ErrUnknownCategory = errors.New("unknown category")
