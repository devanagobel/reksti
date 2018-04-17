package main

import (
	"database/sql"
	"log"
	//"errors"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	Nim string `json:"nim", sql:"student_nim"`
	Name string `json:"name", sql:"student_name"`
	Faculty string `json:"faculty", sql:"student_faculty"`
	Major	string	`json:"major", sql:"student_major"`
}


func (student *Student) getAllStudent() (err error){
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/reksti")
	if err != nil {
		log.Fatalf("cannot open database")
		return
	}
	defer db.Close()

	rows,err := db.Query("SELECT student_nim, student_name, student_faculty, student_major FROM student")
	if err != nil {
		log.Fatalf("error in querying database")
		log.Fatal(err)
		return
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&student.Nim, &student.Name, &student.Faculty, &student.Major)
	}

	if err != nil {
		log.Fatalf("error in scanning databaes")
	}

	return
}