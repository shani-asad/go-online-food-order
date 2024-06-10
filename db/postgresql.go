package db

import (
	"database/sql"
	"log"
	"online-food/model/properties"
	"time"

	_ "github.com/lib/pq"
)

func InitPostgreDB(config properties.PostgreConfig) *sql.DB {
	connStr := config.DatabaseURL

	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(10)
    db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(time.Minute * 30)
	db.SetConnMaxLifetime(time.Hour * 1)

	return db
}