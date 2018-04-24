package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"errors"
)

type Course struct {
	Index 	string 	`json:"course_index"`
	Name	string	`json:"course_name"`
}

func (course *Course) getCourseName() (err error) {

	query := "SELECT course_index, course_name FROM course WHERE course_index = " + "'" + course.Index + "'"

	if course.Index == "" {
		err = errors.New("cannot get Course Data")
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

func (course *Course) getAllCourse() (err error, result []Course) {
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/reksti")
	if err != nil {
		log.Fatalf("cannot open database")
		return
	}
	defer db.Close()

	var data Course

	rows, err := db.Query("SELECT course_index, course_name FROM course ")
	if err != nil {
		log.Fatalf("error in querying database")
		log.Fatal(err)
		return
	}
	defer rows.Close()

	for rows.Next(){
		err = rows.Scan(&data.Index, &data.Name)
		result = append(result,data)
	}

	if err != nil {
		log.Fatalf("error in scanning database")
	}
	return
}