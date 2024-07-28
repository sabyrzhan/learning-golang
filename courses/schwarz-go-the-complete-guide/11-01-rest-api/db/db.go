package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	open, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}

	DB = open
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createEventsTableQuery := `
CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    date_time DATETIME NOT NULL,
    user_id INTEGER NOT NULL
)
`
	_, err := DB.Exec(createEventsTableQuery)
	if err != nil {
		panic(err)
	}
}
