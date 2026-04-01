package main

import "fmt"

func main() {
	// Bidirectional channel: can send and receive.
	ch := make(chan int)

	go func() {
		ch <- 42 // send
	}()

	val := <-ch // receive
	fmt.Println("Received:", val)
}
