package main

import "fmt"

// producer owns the channel it creates.
// It is responsible for sending values and closing the channel.
// It returns a receive‑only channel to enforce that the caller cannot send.
func producer() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // producer closes – it's the owner.
	}()
	return ch
}

// consumer only receives from the channel. It never sends and never closes.
func consumer(ch <-chan int) {
	// range automatically stops when channel is closed.
	for v := range ch {
		fmt.Println("Received:", v)
	}
}

func main() {
	ch := producer() // get a read‑only channel
	consumer(ch)     // consume all values
	// channel is already closed; no need to close again.
}
