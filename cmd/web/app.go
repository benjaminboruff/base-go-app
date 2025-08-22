package main

import (
	"database/sql"
	// "github.com/benjaminboruff/base-go-app/internal/models"
)

// App defines a struct to hold
// applications-wide dependencies and
// application settings.
type App struct {
	Addr     string
	Database *sql.DB
	HTMLDir  string
	DistDir  string
}
