package main

import (
	"database/sql"
	"github.com/benjaminboruff/base-go-app/internal/models"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"

	"github.com/alexedwards/scs/v2"
)

type Env struct {
	users interface {
		All() ([]models.User, error)
		Close()
		CreateTable() error
		Show(int) (models.User, error)
		Seed() error
		Create(models.User) (int64, error)
		Verify(string, string) (bool, error)
	}
}

func main() {

	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	dsn := "./base-go-app.db"
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal(err)
	}

	env := &Env{
		users: models.UserModel{DB: db},
	}
	defer env.users.Close()

	err = env.users.CreateTable()
	if err != nil {
		// If the users table cannot be created
		// then crash, cuz something bad is afoot!
		log.Println("Cannot create users table. Time to die!")
		log.Fatal(err)
	}

	// err = env.users.SeedUsers()
	// if err != nil {
	// 	log.Println(err)
	// }

	app := &App{
		Addr:           ":8080",
		Database:       db,
		Env:            env,
		HTMLDir:        "./ui/html",
		DistDir:        "./ui/dist",
		SessionManager: sessionManager,
	}

	app.RunServer()
}
