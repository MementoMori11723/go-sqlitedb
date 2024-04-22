package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)
	

func main() {
	db, err := sql.Open("sqlite3", "database/data.db")
	if err != nil {
		panic(err)
	}
	log.Println("Connected to the database")
	db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")
	log.Println("Created table users")
	db.Exec("INSERT INTO users (name, age) VALUES ('Alice', 25)")
	log.Println("Inserted user Alice")
	db.Exec("INSERT INTO users (name, age) VALUES ('Bob', 30)")
	log.Println("Inserted user Bob")
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id, age int
		var name string
		rows.Scan(&id, &name, &age)
		fmt.Println(id, name, age)
	}
	defer db.Close()
}