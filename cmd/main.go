package main

import (
	internal "github.com/Edmond-develop/subscription-tracker/internal/config"
	database "github.com/Edmond-develop/subscription-tracker/internal/database"
)

func main() {
	cfg := internal.LoadConfig()
	db := database.ConnectDB(cfg.Database.Name)
	_ = db
}
