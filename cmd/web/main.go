package main

import (
	"log"

	"github.com/benjaminboruff/base-go-app/internal/models"
	"github.com/benjaminboruff/base-go-app/internal/utils"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	dsn := "./base-go-app.db"
	connection := utils.Connect(dsn)
	db := &models.Database{Connection: connection}
	// Defer the closing of the database connection
	defer db.Close()

	err := db.CreateUsersTable()
	if err != nil {
		// If the users table cannot be created
		// then crash, cuz something bad is afoot!
		log.Println("Cannot create users table. Time to die!")
		log.Fatal(err)

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
