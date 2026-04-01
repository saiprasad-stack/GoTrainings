package main

import (
	"fmt"
	"time"
)

func main() {
	// A channel of empty structs uses zero memory – ideal for signaling.
	done := make(chan struct{})

	// Start a worker goroutine.
	go func() {
		fmt.Println("Goroutine working...")
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine done, signaling.")

		// Closing the channel broadcasts to all receivers that work is finished.
		// No data is sent; the close itself is the signal.
		close(done)
	}()

	// Main goroutine waits for the signal.
	// The receive operation blocks until the channel is closed.
	<-done
	fmt.Println("Main received done signal")
}
