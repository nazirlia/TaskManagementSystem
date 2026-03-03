package main

import (
	"fmt"
	"session-10/task/go_generics"
)

func main() {
	// Task 1
	fmt.Println("Sum:", go_generics.Sum([]int{1, 2, 3, 4}))
	fmt.Println("Sum:", go_generics.Sum([]float64{1.5, 2.5, 3.5}))

	// Task 2
	min, max := go_generics.MinMax([]int{100, 8, 20, 5, 10, 15})
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)

	// Task 3
	ints := []int{1, 2, 3, 4}
	doubled := go_generics.Map(ints, func(n int) int { return n * 2 })
	fmt.Println(doubled)

	words := []string{"hello", "world"}
	lengths := go_generics.Map(words, func(s string) int { return len(s) })
	fmt.Println(lengths)

	// Task 4
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := go_generics.Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println(evens)

	names := []string{"Ali", "Aysel", "Tural", "Anar"}
	aNames := go_generics.Filter(names, func(s string) bool { return s[0] == 'A' })
	fmt.Println(aNames)
}
