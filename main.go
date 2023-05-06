package main

import (
	"api-go-rest/database"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	database.ConnectDB()
	fmt.Println("Iniciando o servidor Rest com GO")

	routes.HandleRequest()
}
