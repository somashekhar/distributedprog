package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Student struct {
	RollNumber string `json:"rollNumber"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
}

var Students []Student

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home Page!")
}

func returnAllStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Get All Students!\n")
	json.NewEncoder(w).Encode(Students)
}

func returnSingleStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Get A Single student!\n")
	vars := mux.Vars(r)
	key := vars["RollNumber"]

	for _, student := range Students {
		if student.RollNumber == key {
			json.NewEncoder(w).Encode(student)
		}
	}
}

func addNewStudentDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Add a new student!\n")
	var student Student
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &student)

	Students = append(Students, student)
	json.NewEncoder(w).Encode(student)
	fmt.Println(student)
}

func updateStudentDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Updating an existing student record\n")
	vars := mux.Vars(r)
	key := vars["RollNumber"]

	var newStudent Student
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newStudent)

	for index, existingStudent := range Students {
		if existingStudent.RollNumber == key {
			Students[index] = newStudent
			json.NewEncoder(w).Encode(Students[index])
		}
	}

	//json.NewEncoder(w).Encode(Students)
}

func deleteStudentDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Delete a student!\n")
	vars := mux.Vars(r)
	key := vars["RollNumber"]

	for index, student := range Students {
		if student.RollNumber == key {
			Students = append(Students[:index], Students[index+1:]...)
		}
	}
	fmt.Fprintf(w, "Deleting of student record complete!\n")

}

func handleHttpRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)

	router.HandleFunc("/student", addNewStudentDetails).Methods("POST")
	router.HandleFunc("/student/{RollNumber}", updateStudentDetails).Methods("PUT")
	router.HandleFunc("/student/{RollNumber}", deleteStudentDetails).Methods("DELETE")

	//Adding GET a single student here after DELETE, explore why ?
	router.HandleFunc("/student", returnAllStudents)
	router.HandleFunc("/student/{RollNumber}", returnSingleStudent)

	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	Students = []Student{
		Student{RollNumber: "1", FirstName: "Som", LastName: "B"},
		Student{RollNumber: "2", FirstName: "Amar", LastName: "K"},
	}
	handleHttpRequest()
}
