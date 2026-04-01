package main

import "fmt"

func main() {
	// var ch chan int // nil channel
	fmt.Println("Nil Channels")
	// Sending on a nil channel blocks forever (deadlock)
	// ch <- 10 // fatal error: all goroutines are asleep - deadlock!

	// Receiving on a nil channel also blocks forever
	// <-ch // deadlock

	// Closing a nil channel panics
	// close(ch) // panic: close of nil channel
}
