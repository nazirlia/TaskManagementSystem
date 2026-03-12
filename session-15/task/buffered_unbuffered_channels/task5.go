// package buffered_unbuffered_channels
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	ch <- 10
	ch <- 20
	ch <- 30
	close(ch)

	fmt.Println("Sent values into buffered channel")

	go func() {
		for i := range ch {
			fmt.Println(i)
		}
		fmt.Println("All values received")
	}()
	time.Sleep(time.Second * 1)
}
