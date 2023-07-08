package invoiceheader

import "time"
 //model of ivoice header
type Model struct {
	ID       uint
	Client   string
	CreateAt time.Time
	UpdateAt time.Time
}

type Storage interface {
	
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
