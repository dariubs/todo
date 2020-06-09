package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// DB general object
var DB *sql.DB
var err error

func init() {
	DB, err = sql.Open("sqlite3", "./todo.db")
	if err != nil {
		log.Fatal(err)
	}
	DB.Ping()
}
