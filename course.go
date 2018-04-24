package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"errors"
)

type Course struct {
	Id		int	`json:"course_id"`
	Index 	string 	`json:"course_index"`
	Name	string	`json:"course_name"`
}

func (course *Course) getCourseName() (err error) {

	if course.Id == 0 {
		err = errors.New("cannot get Course Data")
		log.Fatal(err)
	}

	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/reksti" )
	if err != nil {
		log.Fatalf("cannot open database")
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT course_index, course_name FROM course WHERE course_id = ?", course.Id)
	if err != nil {
		log.Fatalf("error in querying database")
		log.Fatal(err)
		return
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&course.Index, &course.Name)
	}

	if err != nil {
		log.Fatal(err)
		log.Fatalf("error in scanning database")
	}

	return

}