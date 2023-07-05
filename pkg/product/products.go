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
