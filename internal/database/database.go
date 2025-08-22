package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/benjaminboruff/base-go-app/internal/models"
	"github.com/benjaminboruff/base-go-app/internal/utils"
	"log"
	"time"
)

type UserModel struct {
	DB *sql.DB
}

var (
	ErrUserNotFound       = errors.New("models: user not found")
	ErrInvalidCredentials = errors.New("models: invalid user credentials")
)

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
	user := models.User{}
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

func (u UserModel) Create(user models.User) (int64, error) {

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

// ***************
// UserModel utils
// ***************

// func comparePasswordAndHash(password, hash string) (bool, error) {
// 	parts := strings.Split(hash, "$")
// 	if len(parts) != 6 {
// 		return false, errors.New("invalid hash format")
// 	}

// 	params := &utils.Argon2Params{}
// 	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &params.Memory, &params.Iterations, &params.Parallelism)
// 	if err != nil {
// 		return false, err
// 	}

// 	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
// 	if err != nil {
// 		return false, err
// 	}

// 	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
// 	if err != nil {
// 		return false, err
// 	}
// 	params.KeyLength = uint32(len(decodedHash))

// 	comparisonHash := argon2.IDKey([]byte(password), salt, params.Iterations, params.Memory, params.Parallelism, params.KeyLength)

// 	return (subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1), nil
// }
