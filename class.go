package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"errors"
)

type Class struct {
	Index 	string	`json:"class_index"`
	Name 	string	`json:"class_name"`
	Course	string	`json:"course_index"`
}

func (class *Class) getClassName() (err error) {

	query := "SELECT class_index, class_name, course_index FROM class WHERE class_index = " + "'" + class.Index + "'"

	if class.Index == "" {
		err = errors.New("cannot get Class Data")
		log.Fatal(err)
	}

	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/reksti" )
	if err != nil {
		log.Fatalf("cannot open database")
		return
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error in querying database")
		log.Fatal(err)
		return
	}
	defer rows.Close()

	if rows.Next(){
		err = rows.Scan(&class.Index, &class.Name, &class.Course)
	}

	if err != nil {
		log.Fatal(err)
		log.Fatalf("error in scanning database")
	}
	return
}

