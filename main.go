package main

import (
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	models.People = []models.Person{
		{Id: 1, Name: "Name 1", History: "history 1"},
		{Id: 2, Name: "Name 2", History: "history 2"},
	}

	fmt.Println("Iniciando o servidor Rest com GO")

	routes.HandleRequest()
}
