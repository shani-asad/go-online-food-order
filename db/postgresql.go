package db

import (
	"database/sql"
	"log"
	"online-food/model/properties"

	_ "github.com/lib/pq"
)

func InitPostgreDB(config properties.PostgreConfig) *sql.DB {
	connStr := config.DatabaseURL

	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
