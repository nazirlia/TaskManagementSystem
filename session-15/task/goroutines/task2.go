package goroutines

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Application started on main thread")

	go func() {
		for i := 'A'; i <= 'E'; i++ {
			fmt.Println(string(i))
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func() {
		for i := 1; i < 6; i++ {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 300)
		}
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Application finished on main thread")
}
