package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Cancelled")
				return
			default:
				fmt.Println(i, "running")
				time.Sleep(1 * time.Second)
			}

		}
	}()

	time.Sleep(3 * time.Second)
	cancel()
	wg.Wait()
}
