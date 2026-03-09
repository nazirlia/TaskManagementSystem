package mocking

import "errors"

type User struct {
	ID   int
	Name string
}

type Database interface {
	GetUserByID(id int) (User, error)
}

func GetUser(db Database, id int) (User, error) {
	if id <= 0 {
		return User{}, errors.New("invalid ID")
	}
	return db.GetUserByID(id)
}
