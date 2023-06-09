package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//get connection. GET connection of the datebase(posdtgres)

func GetConnection()*sql.DB{
	dsn := "postgres://postgres:rivera220@localhost:5432/gocrud?sslmode=disable"
	db, err := sql.Open("postgres",dsn)
	

	if err != nil{
		log.Fatal(err)
	}


	return db
}