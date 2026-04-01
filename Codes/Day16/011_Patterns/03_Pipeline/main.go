package main

import "fmt"

// gen converts a list of numbers into a channel that emits those numbers.
// It returns a receive‑only channel so the caller cannot accidentally send.
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out) // close when done sending
	}()
	return out
}

// sq reads integers from an input channel, squares them, and sends them to an output channel.
// It returns a receive‑only channel.
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in { // reads until in is closed
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// Set up the pipeline.
	// Stage 1: generate numbers 2,3,4
	c := gen(2, 3, 4)

	// Stage 2: square each number
	out := sq(c)

	// Stage 3: consume the squared values.
	// The pipeline stages run concurrently; the main goroutine reads the final output.
	for result := range out {
		fmt.Println(result)
	}
}
