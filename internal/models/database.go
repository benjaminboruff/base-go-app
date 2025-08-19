package models

import (
	"database/sql"
	"log"

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
		"email" VARCHAR(255) UNIQUE,
		"password" TEXT
	);`
	_, err = db.DB.Exec(createTableSQL)

	if err != nil {
		log.Fatal(err)
	}

	u := User{}
	u.Email = "dude@gmail.com"
	u.FirstName = "Dude"
	u.LastName = "Erino"
	u.MiddleName = "Q"
	u.Password = "12345678"

	err = u.Create(db)
	if err != nil {
		log.Fatal(err)
	}

	return err
}
