package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// Model definition

type TodoItem struct {
	ID          int    `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Description string `sql:"size:255;index"`
	Deadline    string
	Progress    int
}

// Database functions

// Create connection to sqlite3 database storing the todos.
func InitDB(databaseName string) *DB {

	// Tries to connect to specified sqlite3 database.
	db, err := gorm.Open("sqlite3", databaseName)
	if err != nil {
		log.Fatal(err)
	}

	// Check connection to database in order to be sure.
	err = db.DB().Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Check if todo table is present. If not, create it.
	db.CreateTable(&TodoItem{})

	return &db
}
