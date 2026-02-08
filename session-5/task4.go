package main

import (
	"fmt"
)

func main() {
	var x int
	var y int
	var z int

	x = 1
	y = 2
	z = 3

	fmt.Printf("Before doubling: x = %d, y = %d, z = %d\n", x, y, z)
	doubleVariadicElements(&x, &y, &z)
	fmt.Printf("After doubling: x = %d, y = %d, z = %d\n", x, y, z)

}

func doubleVariadicElements(numbers ...*int) {
	for _, number := range numbers {
		*number = *number * 2
	}
}
