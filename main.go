package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func seedAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}
	if err := store.CreateAccount(acc); err != nil {
		log.Fatal("failed to create account")
	}
	fmt.Println("account created", acc.Number)
	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "John", "Doe", "password")
}

func main() {
	seed := flag.Bool("seed", false, "seed accounts")
	flag.Parse()

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

	if *seed {
		fmt.Printf("Seeding accounts...\n")
		seedAccounts(store)
	}

	server := NewAPIServer("localhost:8080", store)
	server.Run()
}
