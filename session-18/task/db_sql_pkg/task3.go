package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection successful")
	}
	defer db.Close()

}

func connect() (*sql.DB, error) {
	connStr := "user=student dbname=school password=securepass host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
