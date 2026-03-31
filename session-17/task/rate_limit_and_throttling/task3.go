package main

import (
	"fmt"
	"time"
)

func TaskExecuter() {
	fmt.Println("task executed.")
}

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	count := 0

	for range ticker.C {
		TaskExecuter()
		count++
		if count == 5 {
			fmt.Println("(ticker stopped)")
			break
		}
	}

	time.Sleep(5 * time.Second)
}
