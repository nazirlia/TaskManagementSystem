package main

import "fmt"

func main() {
	var arrInt [5]int
	arrInt = [5]int{10, 20, 30, 40, 50}
	fmt.Println("First element of arrInt:", arrInt[0])
	fmt.Println("Last element of arrInt:", arrInt[len(arrInt)-1])
	var sum int
	for i := 0; i < len(arrInt); i++ {
		sum += arrInt[i]
	}
	fmt.Println("Sum of array's elements:", sum)

	fmt.Print("Reversed Array: [")
	for i := len(arrInt) - 1; i >= 0; i-- {
		fmt.Print(arrInt[i])
		if i != 0 {
			fmt.Print(",")
		}
	}
	fmt.Println("]")

}
