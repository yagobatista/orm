package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Connection struct {
	db *sql.DB
}

var globalConnection *Connection

func Init() {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=orm sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	globalConnection = &Connection{db}
}

func Ping() error {
	return globalConnection.db.Ping()
}

func Shutdown() error {
	return globalConnection.db.Close()
}
