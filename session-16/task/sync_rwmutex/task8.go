package main

import (
	"fmt"
	"sync"
)

func CountReader(counter *int, mu *sync.RWMutex) {
	mu.RLock()
	defer mu.RUnlock()
	fmt.Printf("Reader accessed counter: %d\n", *counter)
}

func CountWriter(counter *int, mu *sync.RWMutex) {
	mu.Lock()
	defer mu.Unlock()
	*counter++
	fmt.Printf("Writer updated counter: %d\n", *counter)
}

func main() {

	var counter int

	var wg sync.WaitGroup
	var mu sync.RWMutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			CountReader(&counter, &mu)
		}()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			CountWriter(&counter, &mu)
		}()
	}

	wg.Wait()

	fmt.Println(counter)

}
