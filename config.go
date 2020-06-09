package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Config model
type Config struct {
	DBPath string `toml:"dbpath"`
}

// Conf general config file
var Conf Config

// LoadConfig file
func LoadConfig() {
	if _, err = toml.DecodeFile("./todo.toml", &Conf); err != nil {
		log.Fatal(err)
	}
}
