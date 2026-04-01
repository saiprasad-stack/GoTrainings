package main

import (
	"fmt"
	"sync"
)

// producer sends a slice of integers into a channel.
func producer(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// worker reads from an input channel, does some work, and writes to an output channel.
// WaitGroup is used to signal when this worker is done.
func worker(id int, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range in {
		// Simulate work (e.g., squaring)
		out <- n * n
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	in := producer(nums...) // produce input values

	// Fan‑out: start multiple workers.
	const numWorkers = 3
	var wg sync.WaitGroup

	// Create a buffered results channel to hold all results.
	results := make(chan int, len(nums))

	// Launch workers.
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, in, results, &wg)
	}

	// Fan‑in: close results channel when all workers are done.
	// This goroutine waits for workers and then closes results.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect all results from the single results channel.
	for r := range results {
		fmt.Println("Result:", r)
	}
}
