package utils

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func GetDatabaseConnection() *sql.DB {
	uri := GetDatabaseConfiguration().Uri

	db, err := sql.Open("postgres", uri)

	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %s", err)
	}

	return db
}

func CreateLinksTable() {
	db := GetDatabaseConnection()

	defer db.Close()

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS links (
			id VARCHAR(12) PRIMARY KEY,
			forward_url VARCHAR(2550) NOT NULL,
			user_id VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)

	if err != nil {
		log.Fatalf("Error creating the links table: %s", err)
	}
}

func DatabaseInit() {
	CreateLinksTable()
}
