package main

import (
	"fmt"
	"log"
)

func main() {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.Exec("UPDATE students SET age = 21 WHERE name = 'Ali'")
	if err != nil {
		err := tx.Rollback()
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Fatal(rbErr)
		}
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transaction successful")
}
