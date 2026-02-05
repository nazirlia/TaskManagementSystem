package main

import "fmt"

func main() {
	var x int
	x = 5

	fmt.Println("Value of x before incrementing by value:", x)
	incrementByValue(x)
	fmt.Println("After incrementing by value:", x)
	fmt.Println("Value of x after incrementing by pointer:", x)
	incrementByPointer(&x)
	fmt.Println("After incrementing by pointer:", x)
}

func incrementByValue(val int) {
	val = val + 1

}

func incrementByPointer(ptr *int) {
	*ptr = *ptr + 1
}
