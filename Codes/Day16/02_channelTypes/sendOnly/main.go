package main

import "fmt"

// sendOnly accepts a send-only channel.
func sendOnly(ch chan<- int, value int) {
	ch <- value
	// <-ch // not allowed here (receive from send-only)
}

func main() {
	ch := make(chan int)

	go sendOnly(ch, 100)

	fmt.Println("Received:", <-ch)
}
