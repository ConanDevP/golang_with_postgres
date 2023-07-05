package invoiceheader

import "time"
 //model of ivoice header
type Model struct {
	ID       uint
	Client   string
	CreateAt time.Time
	UpdateAt time.Time
}