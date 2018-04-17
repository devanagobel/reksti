package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main(){
	port := 8181

	http.HandleFunc("/student/profile", func(w http.ResponseWriter, r *http.Request){
		switch r.Method{
		case "GET":
			handleStudentGetAllProfile(w,r)
		case "PUT":
			break
		}
	})

	log.Printf("Server starting on port %v\n",port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}

func handleStudentGetAllProfile (writer http.ResponseWriter, _ *http.Request ) {
	student := Student{}

	err := student.getAllStudent()
	if err != nil {
		log.Fatalf("error in encoding Student data to JSON")
		writer.WriteHeader(500)
		return
	}

	if student.Name == "" {
		writer.WriteHeader(404)
		return
	} else {
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(&student)
		if err != nil {
			log.Fatalf("error in encoding Student data to JSON")
		}
	}
}