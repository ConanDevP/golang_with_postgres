package main

import (
	"log"

	"github.com/golang_with_postgres/storage"
)

func main() {
	student := storage.Student{
		Name:   "Juan",
		Age:    21,
		Active: true,
	}

	err := storage.CreateStudent(student)

	if err != nil {
		log.Fatal(err)
	}

}
