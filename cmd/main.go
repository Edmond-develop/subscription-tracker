// @title Subscriptions API
// @version 1.0
// @description Простая документация для сервиса подписок

// @host localhost:8080
// @BasePath /
package main

import (
	"github.com/Edmond-develop/subscription-tracker/internal/config"
	"github.com/Edmond-develop/subscription-tracker/internal/controller"
	"github.com/Edmond-develop/subscription-tracker/internal/database"
	"github.com/Edmond-develop/subscription-tracker/internal/repository"
	"github.com/Edmond-develop/subscription-tracker/internal/routes"
	"github.com/Edmond-develop/subscription-tracker/internal/service"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	db, err := database.ConnectDB(cfg)

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	subRepo := repository.NewSubscriptionRepository(db)

	subService := service.NewSubscriptionService(subRepo)

	subHandler := controller.NewSubscriptionHandler(subService)

	router := routes.SetupRoutes(subHandler)

	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
