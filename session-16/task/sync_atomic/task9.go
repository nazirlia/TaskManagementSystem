package sync_atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Increment(counter *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		atomic.AddInt64(counter, 1)
	}
}

func main() {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Increment(&counter, &wg)
	}

	wg.Wait()
	fmt.Println("Final atomic counter value:", counter)
}
