package controllers

import (
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ListPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.People)
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, person := range models.People {
		if strconv.Itoa(person.Id) == id {
			json.NewEncoder(w).Encode(person)
		}
	}
}
