package main

import (
	"log"
)

// Todo model
type Todo struct {
	ID     uint
	Task   string
	Status string // values: active, doing, done, archived
}

// Add todo to database
func (t *Todo) Add() error {
	stmt, err := DB.Prepare("CREATE TABLE IF NOT EXISTS todo (id INTEGER PRIMARY KEY, task TEXT, status TEXT)")
	if err != nil {
		log.Println(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Println(err)
	}
	stmt, err = DB.Prepare("INSERT INTO todo (task, status) VALUES (?, ?)")
	if err != nil {
		log.Println(err)
	}
	_, err = stmt.Exec(t.Task, t.Status)
	if err != nil {
		log.Println(err)
	}
	return nil
}

// ChangeStatus of todo item
func (t *Todo) ChangeStatus(status string) error {
	stmt, err := DB.Prepare("UPDATE todo set status = ? WHERE ID = ?")
	if err != nil {
		log.Println(err)
	}
	_, err = stmt.Exec(status, t.ID)
	if err != nil {
		log.Println(err)
	}
	return nil
}

// Get list of todo items
func Get() ([]Todo, error) {
	rows, err := DB.Query("SELECT id, task, status FROM todo")
	if err != nil {
		log.Println(err)
	}
	var t []Todo

	for rows.Next() {
		ti := Todo{}
		rows.Scan(&ti.ID, &ti.Task, &ti.Status)
		t = append(t, ti)
	}
	return t, nil
}
