package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Student struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	RollNumber int    `json:"rollNumber"`
}

var Students []Student

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home Page !")
	fmt.Println("Endpoint Hit: HomePage")

}

func AllStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Returning list of Students!")
	json.NewEncoder(w).Encode(Students)
}

func handleHttpRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/student", AllStudents)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	Students = []Student{
		Student{FirstName: "Som", LastName: "Bidari", RollNumber: 1},
		Student{FirstName: "Amar", LastName: "K", RollNumber: 2},
	}
	handleHttpRequest()
}
