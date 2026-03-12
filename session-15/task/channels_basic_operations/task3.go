package channels_basic_operations

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- 42
	}()
	fmt.Println(<-ch)
}
