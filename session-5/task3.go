package main

import "fmt"

func main() {
	var x int
	var y int
	x = 10
	y = 20
	fmt.Printf("Before swapping x = %d, y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swapping x = %d, y = %d\n", x, y)
}

func swap(a, b *int) {
	var tempC int
	tempC = *a
	*a = *b
	*b = tempC
}
