package models

import (
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID         int
	FirstName  string
	MiddleName string
	LastName   string
	Email      string
	Password   string
}

func (u *User) Create(db *Database) error {

	stmt := "INSERT INTO users (first_name, middle_name, last_name, email, password) VALUES (?, ?, ?, ?, ?)"

	_, err := db.DB.Exec(stmt, u.FirstName, u.MiddleName, u.LastName, u.Email, u.Password)

	return err
}
