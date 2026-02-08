package main

import "fmt"

func main() {

	var arrInt [10]int
	var sliceInt []int
	arrInt = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceInt = arrInt[:]
	subSliceInt := sliceInt[2:6]
	fmt.Println("Sub-slice:", subSliceInt)
	subSliceInt = append(subSliceInt, 10, 11, 12)
	fmt.Println("Appended slice:", subSliceInt)
	fmt.Println("Original array:", arrInt)

}
