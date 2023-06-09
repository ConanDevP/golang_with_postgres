package storage

import (
	"errors"
	"log"
	"time"
)

type Student struct {
	ID       uint
	Name     string
	Age      uint16
	Active   bool
	CreateAt time.Time
	UpdateAt time.Time
}

//Create once student in the database

func CreateStudent(e Student) error {

	query := `INSERT INTO
		 student (name,age,active)
		VALUES($1,$2,$3)
		 `

	db := GetConnection()

	defer db.Close()

	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(e.Name, e.Age, e.Active)

	if err != nil {
		log.Fatal(err)
	}

	i, _ := row.RowsAffected()

	if i != 1 {
		return errors.New("Insert filled")
	}

	return nil
}
