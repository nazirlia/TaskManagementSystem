package main

import "fmt"

func main() {

	var x int
	x = 10
	memAddressOfX := &x

	fmt.Println("Value of X is", x)
	fmt.Println("Address of X is", memAddressOfX)
	fmt.Println("Value at pointer is", *memAddressOfX)
}
