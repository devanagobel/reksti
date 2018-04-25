package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	//"strconv"
	"time"

)

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/student/{id}", handleStudentGetProfile).Methods("GET")
	router.HandleFunc("/student", handleGetAllStudent).Methods("GET")
	router.HandleFunc("/course/{id}", handleCourseGetName).Methods("GET")
	router.HandleFunc("/course", handleGetAllCourse).Methods("GET")
	router.HandleFunc("/course/class/{id}", handleGetCourseByClass).Methods("GET")
	router.HandleFunc("/class/{id}", handleGetClassName).Methods("GET")
	router.HandleFunc("/class/course/{id}",handleGetClassByCourse).Methods("GET")
	router.HandleFunc("/class", handleGetAllClass).Methods("GET")


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
	studentNIM := vars["id"]

	student := Student{
		Nim: studentNIM,
	}

	err := student.getStudentProfile()
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
	courseIndex := vars["id"]

	course := Course{
		Index: courseIndex,
	}

	err := course.getCourseName()
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

func handleGetClassName (writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")

	vars := mux.Vars(request)
	classIndex := vars["id"]

	class := Class{
		Index: classIndex,
	}

	err := class.getClassName()
	if err != nil {
		log.Fatal(err)
		log.Fatalf("error in encoding Class data to JSON")
		writer.WriteHeader(500)
		return
	}

	if class.Index == "" || class.Name == "" {
		writer.WriteHeader(404)
		return
	} else {
		encoder := json.NewEncoder(writer)
		err = encoder.Encode(&class)

		if err != nil {
			log.Fatalf("error in encoding Class data to JSON")
		}
	}
}

func handleGetCourseByClass (writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-type", "application/json")
	vars := mux.Vars(request)
	classIndex := vars["id"]

	class := Class{
		Index: classIndex,
	}

	err, course := class.getCourseFromClass()
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

func handleGetAllCourse (writer http.ResponseWriter, request *http.Request ) {
	writer.Header().Set("Content-type", "application/json")

	course := Course{}

	err, courses := course.getAllCourse()
	if err != nil {
		log.Fatal(err)
		log.Fatalf("error in encoding Course data to JSON")
		writer.WriteHeader(500)
		return
	}

	if courses[0].Index == "" || courses[0].Name == "" {
		writer.WriteHeader(400)
		return
	} else {
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(&courses)

		if err != nil {
			log.Fatal(err)
			log.Fatalf("error in encoding Course data to JSON")
		}
	}
}

func handleGetAllClass (writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")

	class := Class{}

	err, classes := class.getAllClasses()
	if err != nil {
		log.Fatal(err)
		log.Fatalf("error in encoding Class data to JSON")
		writer.WriteHeader(500)
		return
	}

	if classes[0].Index == "" || classes[0].Name == "" {
		writer.WriteHeader(400)
		return
	} else {
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(&classes)

		if err != nil {
			log.Fatal(err)
			log.Fatalf("error in encoding Class data to JSON")
		}
	}
}

func handleGetClassByCourse (writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-type", "application/json")

	vars := mux.Vars(request)
	courseIndex := vars["id"]

	class := Class{
		Course: courseIndex,
	}

	err, classes := class.getClassForCourse()
	if err != nil {
		log.Fatal(err)
		log.Fatalf("error in encoding Course data to JSON")
		writer.WriteHeader(500)
		return
	}

	if classes[0].Index == "" || classes[0].Name == "" {
		writer.WriteHeader(404)
		return
	} else {
		encoder := json.NewEncoder(writer)
		err = encoder.Encode(&classes)

		if err != nil {
			log.Fatalf("error in encoding Student data to JSON")
		}
	}
}