package main

import (
	"database/sql"
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
