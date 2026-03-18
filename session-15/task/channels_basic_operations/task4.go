package channels_basic_operations

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 5
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("Channel closed")
}
