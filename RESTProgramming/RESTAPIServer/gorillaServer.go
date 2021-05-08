package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Student struct {
	RollNumber int    `json:"rollNumber"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"LastName"`
}

var Students []Student

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home Page!")
}

func AllStudents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Students)
}

// Using custom router gorilla/mux
func handleHttpRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", AllStudents)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Students = []Student{
		Student{FirstName: "Som", LastName: "Bidari", RollNumber: 1},
		Student{FirstName: "Amar", LastName: "K", RollNumber: 2},
	}
	handleHttpRequest()
}
