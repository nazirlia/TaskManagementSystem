// package buffered_unbuffered_channels
package main

func main() {
	ch := make(chan string)

	ch <- "hello"
}
