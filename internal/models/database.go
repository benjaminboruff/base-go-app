package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	DB *sql.DB
}

func (db *Database) InitDB() error {
	var err error

	// Create users table
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"first_name" VARCHAR(255),
		"middle_name" VARCHAR(255),
		"last_name" VARCHAR(255),
		"email" VARCHAR(255) NOT NULL UNIQUE,
		"password" VARCHAR(255) NOT NULL
	);`
	_, err = db.DB.Exec(createTableSQL)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func (db *Database) Seed() error {

	u := User{}
	u.Email = "dude@gmail.com"
	u.FirstName = "Dude"
	u.LastName = "Erino"
	u.MiddleName = "Q"
	u.Password = "12345678"

	err := u.Create(db)
	if err != nil {
		fmt.Println(err)
	}

	return err
}
