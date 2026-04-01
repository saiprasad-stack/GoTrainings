package main

import (
	"fmt"
	"time"
)

// producer generates data and sends it into the channel.
// It owns the channel and closes it when done.
func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Producing", i)
		ch <- i                            // send value (blocks if buffer full)
		time.Sleep(500 * time.Millisecond) // simulate work
	}
	close(ch) // signal that no more values will be sent
}

// consumer reads values from the channel until it is closed.
func consumer(ch <-chan int) {
	// range automatically stops when channel is closed and drained
	for val := range ch {
		fmt.Println("Consuming", val)
		// no need to close a receive-only channel
	}
}

func main() {
	// Create a buffered channel with capacity 2.
	// This allows the producer to send up to 2 values without waiting for the consumer.
	ch := make(chan int, 2)

	// Run producer in a separate goroutine.
	go producer(ch)

	// Run consumer in the main goroutine (also a goroutine).
	// The main goroutine blocks here until the channel is closed and drained.
	consumer(ch)
}
