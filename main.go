package main

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/routes"
)

func main() {
	database.DatabaseConnection()
	fmt.Println("Starting the REST Server with Go")
	routes.HandleRequest()
}
