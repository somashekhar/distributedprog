package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Student struct {
	RollNumber string `json:"rollNumber"`
	FirstName  string `json:"firName"`
	LastName   string `json:"lastName"`
}

var Students []Student

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home Page !")
}

func returnAllStudents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Students)
}

func returnSingleStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["RollNumber"]

	for _, stud := range Students {
		if stud.RollNumber == key {
			json.NewEncoder(w).Encode(stud)
		}
	}
}

func handleHttpRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/students", returnAllStudents)
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
