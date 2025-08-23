package models

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/benjaminboruff/base-go-app/internal/utils"
	"golang.org/x/crypto/argon2"
	"log"
	"time"
)

type User struct {
	ID         int
	FirstName  string
	MiddleName string
	LastName   string
	Email      string
	Password   string
	CreatedAt  time.Time
}

type UserModel struct {
	DB *sql.DB
}

var (
	ErrUserNotFound       = errors.New("models: user not found")
	ErrInvalidCredentials = errors.New("models: invalid user credentials")
)

func (u *User) GeneratePasswordHash(password string) error {
	params := &utils.Argon2Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	salt := make([]byte, params.SaltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return err
	}

	hash := argon2.IDKey([]byte(password), salt, params.Iterations, params.Memory, params.Parallelism, params.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
	u.Password = fmt.Sprintf(format, argon2.Version, params.Memory, params.Iterations, params.Parallelism, b64Salt, b64Hash)
	return nil
}

// *****************
// UserModel methods
// interact with the
// DB directly
// *****************

func (u UserModel) CreateUsersTable() error {
	var err error

	// Create users table
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"first_name" VARCHAR(255),
		"middle_name" VARCHAR(255),
		"last_name" VARCHAR(255),
		"email" VARCHAR(255) NOT NULL UNIQUE,
		"password" VARCHAR(255) NOT NULL,
		"created_at" CURRENT_TIMESTAMP NOT NULL
	);`
	_, err = u.DB.Exec(createTableSQL)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func (u UserModel) SeedUsers() error {

	passwordFromForm := "12345678"
	user := User{}
	user.Email = "dude@gmail.com"
	user.FirstName = "Dude"
	user.LastName = "Erino"
	user.MiddleName = "Q"

	err := user.GeneratePasswordHash(passwordFromForm)
	if err != nil {
		log.Println("Password hashing error!")
	}

	user.CreatedAt = time.Now().UTC()

	id, err := u.Create(user)
	if err != nil {
		return err
	} else {
		log.Printf("The newly created user's id is: %d", id)
		// return nil
	}
	// verify info from form submission
	verified, err := u.VerifyUser("dude@gmail.com", "12345678")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("the user: %s has been verified: %t\n", user.Email, verified)
	}

	return nil
}

// This ensures the connection is closed when the main function exits

func (u UserModel) Close() {

	log.Println("Disconnected from database.")
	err := u.DB.Close()
	if err != nil {
		log.Printf("Error closing DB: %v", err)
	}
}

// Insert a user into the DB
// Return zero and an error if the insert fails
// Return the user ID and nil if the insert succeeds

func (u UserModel) Create(user User) (int64, error) {

	stmt := "INSERT INTO users (first_name, middle_name, last_name, email, password, created_at) VALUES (?, ?, ?, ?, ?, ?)"

	result, err := u.DB.Exec(stmt, user.FirstName, user.MiddleName, user.LastName, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		log.Println("Could not create user. Notify user.Create form POST response.")
		return 0, err
	} else {
		id, _ := result.LastInsertId()
		return id, nil
	}
}

// Given an email address and password
// verify that the user exists and
// has the correct password

func (u UserModel) VerifyUser(email, password string) (bool, error) {

	var id int
	var hashedPassword string

	row := u.DB.QueryRow("SELECT id, password FROM users WHERE email = $1", email)
	err := row.Scan(&id, &hashedPassword)

	if err == sql.ErrNoRows {
		return false, ErrUserNotFound
	} else if err != nil {
		return false, err
	}

	valid, err := utils.ComparePasswordAndHash(password, hashedPassword)

	if valid != true {
		return false, ErrInvalidCredentials
	} else if err != nil {
		return false, err
	}

	return valid, err
}

// Retrieve all user in the
// user table as roows.

func (u UserModel) All() ([]User, error) {
	rows, err := u.DB.Query("SELECT first_name, middle_name, last_name, email FROM users;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.FirstName, &user.MiddleName, &user.LastName, &user.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Given a user id
// Lookup the user
// and return the user
// data if found

func (u UserModel) ShowUser(id int) (User, error) {

	var user User

	row := u.DB.QueryRow("SELECT first_name, middle_name, last_name, email FROM users WHERE id = ?", id)
	err := row.Scan(&user.FirstName, &user.MiddleName, &user.LastName, &user.Email)

	if err == sql.ErrNoRows {
		return user, ErrUserNotFound
	} else if err != nil {
		return user, err
	}

	return user, err
}
