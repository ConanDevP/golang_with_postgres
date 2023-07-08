## Migrar tabla de productos
```
storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)


if err:= serviceProduct.Migrate(); err != nil{
	log.Fatal(err)
}

```


## Migrar tabla de invoiceheader

```
    storaeInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	serviceInvoiceHeader := invoiceheader.NewService(storaeInvoiceHeader)

	if err := serviceInvoiceHeader.Migrate(); err != nil{
		log.Fatal(err)
	}

```

## Migrar tabla de invoiceitem

```
    storaeInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	serviceInvoiceItem := invoiitem.NewService(storaeInvoiceItem)

	if err := serviceInvoiceItem.Migrate(); err != nil{
		log.Fatal(err)
	}


```