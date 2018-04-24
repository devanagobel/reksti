package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"errors"
)

type Student struct {
	Nim 	string 	`json:"student_nim"`
	Name 	string 	`json:"student_name,omitempty"`
	Faculty string 	`json:"student_faculty,omitempty"`
	Major	string	`json:"student_major,omitempty"`
}


func (student *Student) getStudentProfile() (err error){

	query := "SELECT student_nim, student_name, student_faculty, student_major FROM student WHERE student_nim =" + "'" + student.Nim + "'"
	if student.Nim == "" {
		err = errors.New ("cannot get Student Data")
		log.Fatal(err)
	}

	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/reksti")
	if err != nil {
		log.Fatalf("cannot open database")
		return
	}
	defer db.Close()

	rows,err := db.Query(query)
	if err != nil {
		log.Fatalf("error in querying database")
		log.Fatal(err)
		return
	}
	defer rows.Close()

	if rows.Next() {
		log.Print(student.Nim)
		err = rows.Scan(&student.Nim, &student.Name, &student.Faculty, &student.Major)
	}

	if err != nil {
		log.Fatalf("error in scanning databaes")
	}

	return
}

func (student *Student) getAllStudent() (err error, result []Student) {
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/reksti")
	if err != nil {
		log.Fatalf("cannot open database")
		return
	}
	defer db.Close()

	var data Student

	rows,err := db.Query("SELECT student_nim, student_name, student_faculty, student_major FROM student")
	if err != nil {
		log.Fatalf("error in querying database")
		log.Fatal(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.Nim, &data.Name, &data.Faculty, &data.Major)
		result = append(result, data)
	}

	if err != nil {
		log.Fatalf("error in scanning databaes")
	}
	return
}