package service

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-api/db"
	"rest-api/model"

	"github.com/gorilla/mux"
	memdb "github.com/hashicorp/go-memdb"
)

// DB - Global DB var
var DB *memdb.MemDB

// Init methode to save some users into database and set the datasourc
func init() {

	var err error
	DB, err = db.InitDB()
	if err != nil {

	}

	people = append(people, model.Person{Name: "name", Email: "email@e.com", City: "city", Mac: "08:ub:22:t6", Creditcard: "1555-9999-8521-8477"})
	people = append(people, model.Person{Name: "name", Email: "email@e.com", City: "city", Mac: "08:ub:22:t6", Creditcard: "1555-9999-8521-84377"})
	people = append(people, model.Person{Name: "name", Email: "email@e.com", City: "city", Mac: "08:ub:22:t6", Creditcard: "1555-9999-8521-84477"})
	// Create a write transaction
	txn := DB.Txn(true)

	// Insert a new person
	if err := txn.Insert("person", people[0]); err != nil {
		panic(err)
	}
	if err := txn.Insert("person", people[1]); err != nil {
		panic(err)
	}
	if err := txn.Insert("person", people[2]); err != nil {
		panic(err)
	}
	// Commit the transaction
	txn.Commit()

}

// Get all  data from database
func GetPeople(w http.ResponseWriter, r *http.Request) {

	// Create read-only transaction
	txn := DB.Txn(false)
	defer txn.Abort()

	// Lookup by email
	raw, err := txn.Get("person", "id")
	if err != nil {
		panic(err)
	}

	var result []interface{}
	for {
		row := raw.Next()
		if row == nil {
			break
		}

		result = append(result, row)
	}
	// Say hi!
	if raw != nil {
		json.NewEncoder(w).Encode(result)
	}

}

//Get one Person from saved one into database
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	txn := DB.Txn(false)
	defer txn.Abort()

	// Lookup by email
	var id string
	id = params["id"]
	raw, err := txn.First("person", "id", id)
	if err != nil {
		panic(err)
	}
	if raw != nil {
		json.NewEncoder(w).Encode(raw.(model.Person))
	}
}

//Create Person and saved into db
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	// Create a write transaction
	txn := DB.Txn(true)
	// Decode data
	var Person model.Person

	err := json.NewDecoder(r.Body).Decode(&Person)
	log.Println(err)
	if nil != err {
		// Simplified
		return
	}
	// Insert a new person
	if err := txn.Insert("person", Person); err != nil {
		panic(err)
	}
	// Commit the transaction
	txn.Commit()
}

//Delete methode not implementded as is wasn't asked
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	return
}

var people []model.Person
