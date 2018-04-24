package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"time"

)

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/student/{id:[0-9]+}", handleStudentGetProfile).Methods("GET")
	router.HandleFunc("/student", handleGetAllStudent).Methods("GET")
	router.HandleFunc("/course/{id:[0-9]+}", handleCourseGetName).Methods("GET")

	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Server starting on port 8080\n")
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("%v",err)
	}
}

func handleStudentGetProfile (writer http.ResponseWriter, request *http.Request ) {
	writer.Header().Set("Content-type", "application/json")

	vars := mux.Vars(request)
	studentID, err := strconv.ParseInt(vars["id"], 10, 32)

	if err != nil {
		log.Fatalf("Data not found")
		return
	}

	student := Student{
		Id: int(studentID),
	}

	err = student.getStudentProfile()
	if err != nil {
		log.Fatal(err)
		log.Fatalf("error in encoding Student data to JSON")
		writer.WriteHeader(500)
		return
	}

	if student.Name == "" {
		writer.WriteHeader(404)
		return
	} else {
		encoder := json.NewEncoder(writer)
		err = encoder.Encode(&student)

		if err != nil {
			log.Fatalf("error in encoding Student data to JSON")
		}
	}
}

func handleCourseGetName (writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")

	vars := mux.Vars(request)
	courseID, err := strconv.ParseInt(vars["id"], 10, 32)

	if err != nil {
		log.Fatalf("Data not found")
		return
	}

	course := Course{
		Id: int(courseID),
		Index: "XXXXX",
	}

	err = course.getCourseName()
	if err != nil {
		log.Fatal(err)
		log.Fatalf("error in encoding Course data to JSON")
		writer.WriteHeader(500)
		return
	}

	if course.Index == "" || course.Name == "" {
		writer.WriteHeader(404)
		return
	} else {
		encoder := json.NewEncoder(writer)
		err = encoder.Encode(&course)

		if err != nil {
			log.Fatalf("error in encoding Student data to JSON")
		}
	}
}
func handleGetAllStudent (writer http.ResponseWriter, request *http.Request ) {
	writer.Header().Set("Content-type", "application/json")

	student := Student{}

	err , students := student.getAllStudent()
	if err != nil {
		log.Fatal(err)
		log.Fatalf("error in encoding Student data to JSON")
		writer.WriteHeader(500)
		return
	}

	if students[0].Name == "" {
		writer.WriteHeader(404)
		return
	} else {
		encoder := json.NewEncoder(writer)
		err = encoder.Encode(&students)

		if err != nil {
			log.Fatalf("error in encoding Student data to JSON")
		}
	}
}