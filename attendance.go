package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"errors"
	"html"
)

type Attendance struct {
	Course 	string	`json:"course_index"`
	Class	string	`json:"class_name"`
	Student	string	`json:"student_nim"`
}

func (attendance *Attendance) getAllAttendanceData() (err error, result []Attendance){
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/reksti")
	if err != nil {
		log.Fatal(err)
		log.Fatalf("Cannot open database")
		return
	}
	defer db.Close()

	var data Attendance

	rows, err := db.Query("SELECT course_index, class_index, student_nim FROM attendance")
	if err != nil {
		log.Fatal(err)
		log.Fatalf("error in querying database")
		return
	}
	defer rows.Close()

	for rows.Next(){
		err = rows.Scan(&data.Course, &data.Class, &data.Student)
		result = append(result, data)
	}

	if err != nil {
		log.Fatal(err)
		log.Fatalf("error in scanning database")
	}
	return
}

func (attendance *Attendance) getAttendanceByNIM() (err error, result []Attendance){
	query := "SELECT course_index, class_index FROM attendance WHERE student_nim = " + "'" + attendance.Student + "'"

	if attendance.Student == ""{
		err = errors.New("cannot get Attendance Data")
		log.Fatal(err)
	}

	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/reksti" )
	if err != nil {
		log.Fatal(err)
		log.Fatalf("cannot open database")
		return
	}
	defer db.Close()

	var data Attendance

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
		log.Fatalf("error in querying database")
		return
	}
	defer rows.Close()


	for rows.Next(){
		err = rows.Scan(&data.Course, &data.Class)
		result = append(result, data)
	}

	if err != nil {
		log.Fatal(err)
		log.Fatalf("error in scanning database")
	}
	return
}

func (attendance *Attendance) getAttendanceByCourse() (err error, result []Attendance){
	query := "SELECT class_index, student_nim FROM attendance WHERE course_index = " + "'" + attendance.Course + "'"

	if attendance.Course == "" {
		err = errors.New("cannot get Attendance Data")
		log.Fatal(err)
	}

	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/reksti" )
	if err != nil {
		log.Fatal(err)
		log.Fatalf("cannot open Database")
		return
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
		log.Fatalf("error in querying database")
		return
	}
	defer rows.Close()

	var data Attendance

	for rows.Next(){
		err = rows.Scan(&data.Class, &data.Student)
		result = append(result, data)
	}

	if err != nil {
		log.Fatal(err)
		log.Fatalf("error in scanning database")
	}
	return

}

func (attendance *Attendance) insertAttendanceData() (err error){
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/reksti" )
	if err != nil {
		log.Fatal(err)
		log.Fatalf("cannot open database")
		return
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		log.Fatalf("cannot start database Attendance")
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	if attendance.Course != "" {
		stmt, err := db.Prepare("INSERT INTO attendance (course_index, class_index, student_nim) VALUES (?,?,?)")
		if err != nil {
			log.Fatal(err)
			log.Fatalf("error in preparation INSERT query")
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(html.EscapeString(attendance.Course), html.EscapeString(attendance.Class), html.EscapeString(attendance.Student) )
		if err != nil {
			log.Fatal(err)
			log.Fatalf("error in inserting attendance data")
			return err
		}
	}
	return
}
