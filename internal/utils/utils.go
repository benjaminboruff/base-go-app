package utils

import (
	"database/sql"
	"log"
)

// utils
func Connect(dsn string) *sql.DB {
	connection, err := sql.Open("sqlite3", dsn)

	// This will not be a connection error, but a DSN
	// parse error or another initialization error.
	if err != nil {
		log.Println("DSN parse error or another initialization error. Time to die!")
		log.Fatal(err)
	}

	err = connection.Ping()
	if err != nil {
		log.Println("Connection error. Time to die!")
		log.Fatal(err)
	}

	return connection
}
