package goroutines

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Application started on main thread")
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Application exited on main thread")

}
