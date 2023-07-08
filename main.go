package main

import (
	"log"

	"github.com/golang_with_postgres/pkg/invoiceheader"
	"github.com/golang_with_postgres/pkg/invoiitem"
	"github.com/golang_with_postgres/pkg/product"
	"github.com/golang_with_postgres/pkg/storage"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)


	if err:= serviceProduct.Migrate(); err != nil{
		log.Fatal(err)
	}


	storaeInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	serviceInvoiceHeader := invoiceheader.NewService(storaeInvoiceHeader)

	if err := serviceInvoiceHeader.Migrate(); err != nil{
		log.Fatal(err)
	}

	storaeInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	serviceInvoiceItem := invoiitem.NewService(storaeInvoiceItem)

	if err := serviceInvoiceItem.Migrate(); err != nil{
		log.Fatal(err)
	}


}


