package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)
	

func main() {
	db, err := sql.Open("sqlite3", "sqlite:data.db")
	if err != nil {
		panic(err)
	}
	log.Println("Connected to the database")
	defer db.Close()
}