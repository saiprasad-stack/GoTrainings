package main

import "fmt"

// receiveOnly accepts a receive-only channel.
func receiveOnly(ch <-chan int) {
	val := <-ch
	// ch <- 200 // not allowed here (send to receive-only)
	fmt.Println("Received inside receiveOnly:", val)
}

func main() {
	ch := make(chan int)

	go func() {
		ch <- 300
	}()

	receiveOnly(ch)
}
