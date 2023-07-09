package storage

import (
	"database/sql"
	_"github.com/lib/pq"
	"fmt"
	"log"
	"sync"
)

var (
	db *sql.DB
	once sync.Once
)

func NewPostgresDB(){
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres","postgres://user:password@localhost/dbname?sslmode=disable")

		if err != nil{
			log.Fatalf("Algo salio mal %v",err)
		}

		if err = db.Ping(); err!= nil{
			log.Fatalf("Algo salio mal %v",err)

		}
		fmt.Println("todo chil")


	})
}

func Pool()*sql.DB{
	return db
}
