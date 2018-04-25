package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"time"

)

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/attendance", func (writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer,request, "public/attendance.html")
	})
	router.HandleFunc("/API/student/{id}", handleStudentGetProfile).Methods("GET")
	router.HandleFunc("/API/student", handleGetAllStudent).Methods("GET")

	router.HandleFunc("/API/course/{id}", handleCourseGetName).Methods("GET")
	router.HandleFunc("/API/course", handleGetAllCourse).Methods("GET")
	router.HandleFunc("/API/course/class/{id}", handleGetCourseByClass).Methods("GET")

	router.HandleFunc("/API/class/{id}", handleGetClassName).Methods("GET")
	router.HandleFunc("/API/class/course/{id}",handleGetClassByCourse).Methods("GET")
	router.HandleFunc("/API/class", handleGetAllClass).Methods("GET")

	router.HandleFunc("/API/attendance/course/{id}", handleGetAttendanceByCourse).Methods("GET")
	router.HandleFunc("/API/attendance/student/{id}", handleGetAttendanceByStudent).Methods("GET")
	router.HandleFunc("/API/attendance", handleGetAllAttendance).Methods("GET")
	router.HandleFunc("/API/attendance", handleAttendancePOST).Methods("POST")

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

func handleGetAttendanceByStudent (writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-type", "application/json")

	vars := mux.Vars(request)
	attendanceStudent := vars["id"]

	attendance := Attendance{
		Student: attendanceStudent,
	}

	err, attendances := attendance.getAttendanceByNIM()
	if err != nil {
		log.Fatal(err)
		log.Fatalf("error in encoding Attendance data to JSON")
		writer.WriteHeader(500)
		return
	}

	if attendances[0].Course == "" || attendances[0].Class == "" {
		writer.WriteHeader(404)
		return
	} else {
		encoder := json.NewEncoder(writer)
		err = encoder.Encode(&attendances)

		if err != nil {
			log.Fatalf("error in encoding Attendance data to JSON")
		}
	}
}

func handleGetAttendanceByCourse (writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-type", "application/json")

	vars := mux.Vars(request)
	attendanceCourse := vars["id"]

	attendance := Attendance{
		Course: attendanceCourse,
	}

	err, attendances := attendance.getAttendanceByCourse()
	if err != nil {
		log.Fatal(err)
		log.Fatalf("error in encoding Attendance data to JSON")
		writer.WriteHeader(500)
		return
	}

	if attendances[0].Class == "" || attendances[0].Student == "" {
		writer.WriteHeader(404)
		return
	} else {
		encoder := json.NewEncoder(writer)
		err = encoder.Encode(&attendances)

		if err != nil {
			log.Fatalf("error in encoding Attendance data to JSON")
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

func handleGetAllAttendance (writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-type", "application/json")

	attendance := Attendance{}

	err, attendances := attendance.getAllAttendanceData()
	if err != nil {
		log.Fatal(err)
		log.Fatalf("error in encoding Attendance data to JSON")
		writer.WriteHeader(500)
		return
	}

	if attendances[0].Student == "" || attendances[0].Course == "" {
		writer.WriteHeader(400)
		return
	} else {
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(&attendances)

		if err != nil {
			log.Fatal(err)
			log.Fatalf("error in encoding Attendance data to JSON")
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

func handleAttendancePOST (writer http.ResponseWriter, request *http.Request) {

	attendance := Attendance{}

	err := json.NewDecoder(request.Body).Decode(&attendance)
	if err != nil {
		log.Fatal(err)
	}

	attendance.insertAttendanceData()

	userAttendance, err := json.Marshal(attendance)
	if err != nil {
		log.Fatal(err)
	}
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(userAttendance)
	//
	//if request.Body != nil {
	//	defer request.Body.Close()
	//}
	//
	//
	//decoder := json.NewDecoder(request.Body)
	//
	//attendance := Attendance{}
	//
	//err := decoder.Decode(attendance)
	//if err != nil {
	//	log.Fatal(err)
	//	log.Fatalf("Cannot decode attendance data")
	//	writer.WriteHeader(500)
	//	return
	//}
	//
	//data, err := json.Marshal(attendance)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//writer.Write(data)
	//
	//err = attendance.insertAttendanceData()
	//
	//if err != nil {
	//	writer.WriteHeader(500)
	//	return
	//}
	//
	//attendanceReturn := Attendance {
	//	Course: attendance.Course,
	//	Class: attendance.Class,
	//	Student: attendance.Student,
	//}
	//
	//encoder := json.NewEncoder(writer)
	//encoder.Encode(attendanceReturn)

}