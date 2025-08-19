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
	var err error
	_, err = db.DB.Exec("INSERT INTO users (first_name, middle_name, last_name, email, password) VALUES (?, ?, ?, ?, ?)", u.FirstName, u.MiddleName, u.LastName, u.Email, u.Password)
	return err
}
