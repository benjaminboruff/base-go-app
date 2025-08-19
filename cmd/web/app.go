package main

import (
	"github.com/benjaminboruff/base-go-app/internal/models"
)

// App defines a struct to hold
// applications-wide dependencies and
// application settings.
type App struct {
	Addr     string
	Database *models.Database
	HTMLDir  string
	DistDir  string
}
