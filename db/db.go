package db

import (
	"database/sql"
	"fmt"

	// We need this import, but we will not use it directly
	// _ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "./data/event-navigator-backend.db") //"sqlite3"
	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	email VARCHAR(255) UNIQUE NOT NULL,
    	password VARCHAR(255) NOT NULL,
    	activated int(10) NOT NULL DEFAULT 1
	)`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table." + err.Error())
	}

	createEventTable := `CREATE TABLE IF NOT EXISTS events (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	title VARCHAR(255) NOT NULL,
    	description TEXT NOT NULL,
    	location VARCHAR(255) NOT NULL,
    	organizer VARCHAR(255) NOT NULL,
    	datetime DATETIME NOT NULL,
    	calendar_week int NOT NULL,
    	year int NOT NULL,
    	creator_id int NOT NULL,
    	 
    	FOREIGN KEY (creator_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createEventTable)
	if err != nil {
		panic("Could not create events table." + err.Error())
	}

	fmt.Println("Created events table.")

	createRegistrationsTable := `CREATE TABLE IF NOT EXISTS registrations (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	event_id INTEGER NOT NULL,
    	user_id INTEGER NOT NULL,
    	FOREIGN KEY (event_id) REFERENCES events(id),
    	FOREIGN KEY (user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		panic("Could not create registrations table." + err.Error())
	}
}
