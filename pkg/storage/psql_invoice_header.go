package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlMigrateInvoiceHeader= `CREATE TABLE IF NOT EXISTS invoice_header(
		id SERIAL NOT NULL,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		update_at TIMESTAMP,
		CONSTRAINT invoice_header_id_pk PRIMARY KEY (id)
	)`
)
type PsqlInvoiceHeader struct {
	db *sql.DB
}





func NewPsqlInvoiceHeader(db *sql.DB)*PsqlInvoiceHeader{
	return &PsqlInvoiceHeader{db}
}

func (p *PsqlInvoiceHeader)Migrate()error{
	stmt, err:= p.db.Prepare(psqlMigrateInvoiceHeader)

	if err != nil{
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil{
		return err
	}

	fmt.Println("migracion de invoive header ejecutada")
	return nil
}