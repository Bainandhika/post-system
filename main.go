package main

import (
	"fmt"
	"log"
	"post-system/app/configs"
	"post-system/app/connections"
	"post-system/app/logging"
	"post-system/app/routes"
)

func main() {
	configs.InitConfig()
	appConfig := configs.App

	logger := logging.LoggerConfig{LogPath: appConfig.LogPath}
	logger.InitLogger()

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
	router.Run(fmt.Sprintf("%s:%d", appConfig.Host, appConfig.Port))
}
