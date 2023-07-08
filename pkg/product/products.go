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
	/* Create(*Model) error
	Update(*Model) error
	GetAll() (Model, error)
	GetById(uint) (*Model, error)
	Delete(uint) error */
	Migrate()error
}
//use to manage bussiness logic
type Service struct{
	storae Storage
}

func NewService(s Storage)*Service{
	return &Service{s}
}

//migrate is use for migrate product to the database
func (s *Service)Migrate()error{
	return s.storae.Migrate()	
}









