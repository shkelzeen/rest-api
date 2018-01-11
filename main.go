package main

import (
	"log"
	"net/http"
	"rest-api/service"
	"time"

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

type AppContext struct {
	Render  *render.Render
	Version string
	Env     string
	Port    string
	DB      DataStorer
}

// Healthcheck will store information about its name and version
type Healthcheck struct {
	AppName string `json:"appName"`
	Version string `json:"version"`
}

// Status is a custom response object we pass around the system and send back to the customer
// 404: Not found
// 500: Internal Server Error
type Status struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// CreateContextForTestSetup initialises an application context struct
// for testing purposes
func CreateContextForTestSetup() AppContext {
	testVersion := "0.0.0"
	db := CreateMockDatabase()
	ctx := AppContext{
		Render:  render.New(),
		Version: testVersion,
		Env:     local,
		Port:    "3001",
		DB:      db,
	}
	return ctx
}

// CreateMockDatabase initialises a database for test purposes
func CreateMockDatabase() *MockDB {
	list := make(map[int]User)
	dt, _ := time.Parse(time.RFC3339, "1985-12-31T00:00:00Z")
	list[0] = User{0, "John", "Doe", dt, "London"}
	dt, _ = time.Parse(time.RFC3339, "1992-01-01T00:00:00Z")
	list[1] = User{1, "Jane", "Doe", dt, "Milton Keynes"}
	return &MockDB{list, 1}
}
