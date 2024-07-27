package main

import (
	"log"
	"post-system/app/connections"
	"post-system/app/routes"
)

func main() {
	dbInstance, err := connections.InitDB()
	if err != nil {
        log.Fatal(err)
    }

	dbConnection, err := dbInstance.DB()
	if err != nil {
        log.Fatal(err)
    }
	defer dbConnection.Close()

	router := routes.SetUpRoutes(dbInstance)
	router.Run(":8080")
}
