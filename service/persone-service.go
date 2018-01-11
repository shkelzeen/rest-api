package service

import (
	"encoding/json"
	"net/http"
	"rest-api/model"

	"github.com/gorilla/mux"
)

func init() {

	people = append(people, model.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &model.Address{City: "City X", State: "State X"}})
	people = append(people, model.Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &model.Address{City: "City Z", State: "State Y"}})
	people = append(people, model.Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
		json.NewEncoder(w).Encode(model.Person{})

	}
}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

var people []model.Person
