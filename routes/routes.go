package routes

import (
	"api-go-rest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/people", controllers.ListPerson).Methods("GET")
	r.HandleFunc("/api/people/{id}", controllers.GetPeople).Methods("GET")
	r.HandleFunc("/api/people", controllers.CreatePerson).Methods("POST")
	r.HandleFunc("/api/people/{id}", controllers.RemovePerson).Methods("DELETE")
	r.HandleFunc("/api/people/{id}", controllers.EditPerson).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}
