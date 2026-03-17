package sync_waitgroup

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	sums := make(chan int)
	var result int

	integers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	wg.Add(2)

	go func() {
		defer wg.Done()
		var sum int
		for i := 0; i < len(integers)/2; i++ {
			sum += integers[i]
		}
		sums <- sum
	}()

	go func() {
		defer wg.Done()
		var sum int
		for i := len(integers) / 2; i < len(integers); i++ {
			sum += integers[i]
		}
		sums <- sum
	}()

	go func() {
		wg.Wait()
		close(sums)
	}()

	for sum := range sums {
		result += sum
	}

	fmt.Printf("Sum equals to %d", result)
}
