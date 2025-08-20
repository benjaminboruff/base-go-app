package main

import (
	"database/sql"
	"log"

	"github.com/benjaminboruff/base-go-app/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	dsn := "./base-go-app.db"

	db := connect(dsn, &models.Database{})
	err := db.InitDB()
	if err != nil {
		log.Fatal((err))
	}

	err = db.Seed()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	app := &App{
		Addr:     ":8080",
		Database: db,
		HTMLDir:  "./ui/html",
		DistDir:  "./ui/dist",
	}

	app.RunServer()
}

func connect(dsn string, new_db *models.Database) *models.Database {
	db, err := sql.Open("sqlite3", dsn)

	if err != nil {
		log.Fatal()
	}

	new_db.DB = db

	return new_db
}
