package main

import (
	"database/sql"
	"log"
)

func init() {
	LoadConfig()

	DB, err = sql.Open("sqlite3", Conf.DBPath)
	if err != nil {
		log.Fatal(err)
	}
	DB.Ping()
}
