package select_statement

import (
	"fmt"
	"time"
)

func WaitWithTimeout(ch chan string) string {
	select {
	case res := <-ch:
		return res
	case <-time.After(3 * time.Second):
		return "timeout"
	}

}

func main() {

	ch := make(chan string)

	go func() {
		ch <- "something"
	}()
	result := WaitWithTimeout(ch)

	fmt.Println(result)
}
