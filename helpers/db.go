package helpers

import (
	"log"

	"github.com/asdine/storm/v3"
)

var db *storm.DB
var err error

// OpenDB ...
func OpenDB() *storm.DB {
	db, err = storm.Open("mockup.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// CloseDB ...
func CloseDB() {
	db.Close()
}
