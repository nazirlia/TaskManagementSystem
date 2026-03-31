package main

import (
	"fmt"
	"sync"
	"time"
)

func JobExecuter() {
	fmt.Println("Job started")
	time.Sleep(1 * time.Second)
}

func main() {
	jobLimit := make(chan int, 2)
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			jobLimit <- 1
			JobExecuter()
			<-jobLimit
		}()
	}
	wg.Wait()
}
