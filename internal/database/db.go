package internal

import (
	"database/sql"
	"log"
)

func ConnectDB(url string) *sql.DB {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}
	return db
}
