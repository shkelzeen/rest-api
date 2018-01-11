package main

import (
	"log"
	"net/http"
	"rest-api/service"

	"github.com/gorilla/mux"
)

// our main function
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/people", service.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", service.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", service.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", service.DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
