package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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
		fmt.Printf("ID: %d, name: %s, age: %d\n", id, name, age)
	}
	if err = result.Err(); err != nil {
		log.Fatal(err)
	}

}
