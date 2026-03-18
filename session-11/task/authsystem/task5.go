package authsystem

import (
	"fmt"
)

type UserNotFoundError struct {
	Username string
}

type InvalidPasswordError struct {
	Username string
}

func (e UserNotFoundError) Error() string {
	return fmt.Sprintf("user %s not found", e.Username)
}

func (e InvalidPasswordError) Error() string {
	return fmt.Sprintf("invalid password for user %s", e.Username)
}

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func Authenticate(username, password string) error {
	passValue, ok := users[username]
	if !ok {
		return UserNotFoundError{Username: username}
	}

	if passValue != password {
		return InvalidPasswordError{Username: username}
	}

	return nil
}
