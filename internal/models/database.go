package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Database struct {
	Connection *sql.DB
}

// General DB methods
func (db *Database) CreateUsersTable() error {
	var err error

	// Create users table
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"first_name" VARCHAR(255),
		"middle_name" VARCHAR(255),
		"last_name" VARCHAR(255),
		"email" VARCHAR(255) NOT NULL UNIQUE,
		"password" VARCHAR(255) NOT NULL,
		"created_at" CURRENT_TIMESTAMP NOT NULL
	);`
	_, err = db.Connection.Exec(createTableSQL)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func (db *Database) Seed() error {

	user := User{}
	user.Email = "dude@gmail.com"
	user.FirstName = "Dude"
	user.LastName = "Erino"
	user.MiddleName = "Q"
	user.Password = "12345678"
	user.CreatedAt = time.Now().UTC()

	id, err := user.Create(db)
	if err != nil {
		return err
	} else {
		log.Printf("The newly created user's id is: %d", id)
		return nil
	}
}

// This ensures the connection is closed when the main function exits
func (db *Database) Close() {

	log.Println("Disconnected from database.")
	err := db.Connection.Close()
	if err != nil {
		log.Printf("Error closing DB: %v", err)
	}
}

// User DB methods
//
// CREATE
//
// Insert a user into the DB
// Return zero and an error if the insert fails
// Return the user ID and nil if the insert succeeds

func (u *User) Create(db *Database) (int64, error) {

	stmt := "INSERT INTO users (first_name, middle_name, last_name, email, password, created_at) VALUES (?, ?, ?, ?, ?, ?)"

	result, err := db.Connection.Exec(stmt, u.FirstName, u.MiddleName, u.LastName, u.Email, u.Password, u.CreatedAt)
	if err != nil {
		log.Println("Could not create user. Notify user.Create form POST response.")
		return 0, err
	} else {
		id, _ := result.LastInsertId()
		return id, nil
	}
}
