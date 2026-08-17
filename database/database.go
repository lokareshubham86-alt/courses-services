package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectToDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "course.db")

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to the database successfully")
}

func CreateTable() {
	query := `CREATE TABLE IF NOT EXISTS courses
	 (
	id INTEGER NOT NULL,
	name STRING NOT NULL,
	price INTEGER NOT NULL
	);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table created successfully")
}
