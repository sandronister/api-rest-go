package controllers

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ListPerson(w http.ResponseWriter, r *http.Request) {
	var list []models.Person
	database.DB.Find(&list)
	json.NewEncoder(w).Encode(list)
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var person models.Person
	database.DB.First(&person, id)
	json.NewEncoder(w).Encode(person)
}
