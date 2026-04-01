package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) // unbuffered, capacity 0

	// This send would block forever if not received,
	// so we launch a receiver goroutine.
	go func() {
		time.Sleep(1 * time.Second)
		val := <-ch
		fmt.Println("Received:", val)
	}()

	fmt.Println("Sending 10...")
	ch <- 10 // blocks until receiver is ready
	fmt.Println("Send completed")
}
