package main

import "fmt"

func main() {
	newSlice := make([]int, 3, 5)
	fmt.Printf("Length: %d, Capacity: %d\n", len(newSlice), cap(newSlice))

	initialLen := len(newSlice)
	initialCap := cap(newSlice)

	for i := initialLen; i <= initialCap; i++ {
		newSlice = append(newSlice, i)
		fmt.Printf("Length: %d, Capacity: %d\n", len(newSlice), cap(newSlice))
	}
}
