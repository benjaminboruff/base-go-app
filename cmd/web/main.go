package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/benjaminboruff/base-go-app/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	dsn := "./base-go-app.db"
	db := connect(dsn, &models.Database{})
	db.InitDB()

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

	fmt.Println("Am I here?")
	new_db.DB = db

	return new_db
}
