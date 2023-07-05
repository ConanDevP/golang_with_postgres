package invoiitem

import "time"

//model of invoice item
type Model struct {
	ID              uint
	InvoiceheaderID uint
	ProductID       uint
	CreateAt        time.Time
	UpdateAt        time.Time
}
