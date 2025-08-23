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
	Env      *Env
	HTMLDir  string
	DistDir  string
}
