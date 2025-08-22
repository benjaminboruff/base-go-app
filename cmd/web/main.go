package main

import (
	"database/sql"
	"github.com/benjaminboruff/base-go-app/internal/models"
	"log"
	// "github.com/benjaminboruff/base-go-app/internal/utils"
	_ "github.com/mattn/go-sqlite3"
)

type Env struct {
	users models.UserModel
}

func main() {

	dsn := "./base-go-app.db"
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal(err)
	}

	env := &Env{
		users: models.UserModel{DB: db},
	}
	// connection := utils.Connect(dsn)
	// db := &models.Database{Connection: connection}
	// Defer the closing of the database connection
	defer env.users.Close()

	err = env.users.CreateUsersTable()
	if err != nil {
		// If the users table cannot be created
		// then crash, cuz something bad is afoot!
		log.Println("Cannot create users table. Time to die!")
		log.Fatal(err)
	}

	// err = db.Seed()
	// if err != nil {
	// 	log.Println(err)
	// }

	app := &App{
		Addr:     ":8080",
		Database: db,
		HTMLDir:  "./ui/html",
		DistDir:  "./ui/dist",
	}

	app.RunServer()
}
