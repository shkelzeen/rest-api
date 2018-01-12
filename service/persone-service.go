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

func init() {

	var err error
	DB, err = db.InitDB()
	if err != nil {

	}
	people = append(people, model.Person{ID: "1", Firstname: "John", Lastname: "Doe"})
	people = append(people, model.Person{ID: "2", Firstname: "Koko", Lastname: "Doe"})
	people = append(people, model.Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

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
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	txn := DB.Txn(true)
	defer txn.Abort()

	// Lookup by email
	var id string
	id = params["id"]
	var result int

	err := txn.DeleteAll("person", "id", id)
	if err != nil {
		panic(err)
	}
	//log.Println(result)
}

var people []model.Person
