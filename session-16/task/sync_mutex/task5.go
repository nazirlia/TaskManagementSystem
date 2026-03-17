//package sync_mutex

package main

import (
	"fmt"
	"sync"
)

func Increment(counter *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mu.Lock()
		*counter++
		mu.Unlock()
	}
}

func main() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Increment(&counter, &mu, &wg)
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
