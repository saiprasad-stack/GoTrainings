package main

import "fmt"

func main() {
	// Create an unbuffered channel of strings.
	ch := make(chan string)

	// Start a goroutine that sends a message.
	go func() {
		ch <- "Hello from goroutine!"
	}()

	// Receive the message (this blocks until the send completes).
	msg := <-ch
	fmt.Println(msg)
}
