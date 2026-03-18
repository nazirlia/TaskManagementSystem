package sync_waitgroup

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1 starting")
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine 1 finished")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2 starting")
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine 2 finished")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 3 starting")
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine 3 finished")
	}()

	wg.Wait()
}
