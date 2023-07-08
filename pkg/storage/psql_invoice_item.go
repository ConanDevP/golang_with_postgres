package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlMigraInvoiceItem= `CREATE TABLE IF NOT EXISTS invoice_items(
		id SERIAL NOT NULL,
		invoice_header_id INT NOT NULL,
		products_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		update_at TIMESTAMP,
		CONSTRAINT invoice_items_id_pk PRIMARY KEY (id),
		CONSTRAINT invoice_item_invoice_header_id_fk FOREIGN KEY 
		(invoice_header_id) REFERENCES invoice_header (id) ON UPDATE RESTRICT ON
		DELETE RESTRICT,
		CONSTRAINT invoice_item_product_id_fk FOREIGN KEY 
		(products_id) REFERENCES products (id) ON UPDATE RESTRICT ON
		DELETE RESTRICT
	)` 
)
type PsqlInvoiceItem struct {
	db *sql.DB
}





func NewPsqlInvoiceItem(db *sql.DB)*PsqlInvoiceItem{
	return &PsqlInvoiceItem{db}
}

func (p *PsqlInvoiceItem)Migrate()error{
	stmt, err:= p.db.Prepare(psqlMigraInvoiceItem)

	if err != nil{
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil{
		return err
	}

	fmt.Println("migracion de invoice item  ejecutada")
	return nil
}