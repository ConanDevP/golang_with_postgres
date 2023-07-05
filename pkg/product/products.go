package product

import "time"

//model of products
type Model struct {
	ID          uint
	Name        string
	Observation string
	Price       int
	CreateAt    time.Time
	UpdateAt    time.Time
}

type Models []*Model

type Storage interface {
	Create(*Model) error
	Update(*Model) error
	GetAll() (Model, error)
	GetById(uint) (*Model, error)
	Delete(uint) error
}
