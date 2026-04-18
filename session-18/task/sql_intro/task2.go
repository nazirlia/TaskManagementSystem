package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS students  (id SERIAL PRIMARY KEY, name TEXT, age INT);")
	if err != nil {
		log.Fatal(err)
	}
	// Populate Data
	_, err = db.Exec("INSERT INTO students (name, age) VALUES ('Ali', 20), ('Geray', 22), ('Medina', 21);")
	if err != nil {
		log.Fatal(err)
	}
	// Read
	result, err := db.Query("Select * from students")
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()
	for result.Next() {
		var id int
		var name string
		var age int
		err := result.Scan(
			&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name, age)
	}
	if err = result.Err(); err != nil {
		log.Fatal(err)
	}

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
