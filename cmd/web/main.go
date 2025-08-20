package main

import (
	"database/sql"
	"log"

	"github.com/benjaminboruff/base-go-app/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	dsn := "./base-go-app.db"

	db := connect(dsn)

	// Defer the closing of the database connection
	defer db.Close()

	err := db.CreateUsersTable()
	if err != nil {
		log.Fatal((err))
	}

	err = db.Seed()
	if err != nil {
		log.Println(err)
	}

	app := &App{
		Addr:     ":8080",
		Database: db,
		HTMLDir:  "./ui/html",
		DistDir:  "./ui/dist",
	}

	app.RunServer()
}

func connect(dsn string) *models.Database {
	db, err := sql.Open("sqlite3", dsn)

	// This will not be a connection error, but a DSN
	// parse error or another initialization error.
	if err != nil {
		log.Fatal()
	}

	return &models.Database{db}
}
