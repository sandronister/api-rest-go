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
	r.HandleFunc("/api/people", controllers.ListPeople).Methods("GET")
	r.HandleFunc("/api/people/{id}", controllers.GetPeople).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
