package main

import (
	"errors"
	"fmt"
	"session-11/task/authsystem"
	"session-11/task/custerr"
	"session-11/task/wraperr"
)

func main() {
	// Task 1
	a := 10.0
	b := 2.5
	result, err := custerr.Divide(a, b)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Result: %.2f\n", result)

	// Task 2

	a2 := 20.0
	b2 := 0.0
	result2, err2 := custerr.DivideBy(a2, b2)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(result2)

	// Task 3
	f, err3 := wraperr.OpenFile("go.moad")
	if err3 != nil {
		fmt.Println("Wrapped Error:", err3)
		fmt.Println("Unwrapped Error:", errors.Unwrap(err3))
		return
	}
	f.Close()

	// Task 4
	err4 := wraperr.OpenFile2("go.maaod")
	if err4 != nil {
		fmt.Println(err4)
	}

	// Task 5
	err5 := authsystem.Authenticate("user1", "password1")
	if err5 != nil {
		var unf authsystem.UserNotFoundError
		var ip authsystem.InvalidPasswordError

		switch {
		case errors.As(err, &unf):
			fmt.Println("Login failed:", unf)
		case errors.As(err, &ip):
			fmt.Println("Login failed:", ip)
		default:
			fmt.Println("Login failed:", err)
		}
		return
	}
	fmt.Println("Login succeeded")
}
