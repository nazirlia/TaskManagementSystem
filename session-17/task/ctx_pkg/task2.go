package main

import (
	"context"
	"fmt"
	"time"
)

func someTask(ctx context.Context) {
	fmt.Println("Task started to work")
	select {
	case <-ctx.Done():
		fmt.Println("Cancelled")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Task finished to work")

	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	someTask(ctx)
}
