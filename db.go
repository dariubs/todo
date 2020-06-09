package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// DB general object
var DB *sql.DB
var err error
