package main

import (
	config "github.com/Edmond-develop/subscription-tracker/internal/config"
	database "github.com/Edmond-develop/subscription-tracker/internal/database"
	internal "github.com/Edmond-develop/subscription-tracker/internal/routes"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	db := database.ConnectDB(cfg)
	router := internal.SetupRoutes(db)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
