package database

import (
	"database/sql"
	"fmt"
	config "github.com/Edmond-develop/subscription-tracker/internal/config"
	_ "github.com/lib/pq"
	"log"
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {
	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.Name)
	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}
	return db, nil
}
