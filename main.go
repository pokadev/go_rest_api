package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	// Replace the connection string with your database connection string
	db, err := sql.Open("postgres", "postgres://user:123456@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	store, err := NewPostgresStore(db)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer("localhost:8080", store)
	server.Run()
}
