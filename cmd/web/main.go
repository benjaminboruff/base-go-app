package main

import (
	"database/sql"
	"github.com/benjaminboruff/base-go-app/internal/models"
	_ "github.com/mattn/go-sqlite3"
	"log"
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
	defer env.users.Close()

	err = env.users.CreateUsersTable()
	if err != nil {
		// If the users table cannot be created
		// then crash, cuz something bad is afoot!
		log.Println("Cannot create users table. Time to die!")
		log.Fatal(err)
	}

	err = env.users.SeedUsers()
	if err != nil {
		log.Println(err)
	}

	allUsers, _ := env.users.All()

	log.Println(allUsers)

	app := &App{
		Addr:     ":8080",
		Database: db,
		HTMLDir:  "./ui/html",
		DistDir:  "./ui/dist",
	}

	app.RunServer()
}
