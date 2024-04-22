package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)
	

func main() {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		panic(err)
	}
	log.Println("Connected to the database")
	db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")
	log.Println("Created table users")
	defer db.Close()
}