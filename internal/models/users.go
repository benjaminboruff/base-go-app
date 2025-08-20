package models

import "time"

type User struct {
	ID         int
	FirstName  string
	MiddleName string
	LastName   string
	Email      string
	Password   string
	CreatedAt  time.Time
}
