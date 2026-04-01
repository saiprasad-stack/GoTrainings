package main

import (
	"fmt"
	"sync"
	"time"
)

// worker processes jobs from the jobs channel and sends results to the results channel.
// The WaitGroup is used to signal when this worker has finished.
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	// Ensure Done is called even if the worker panics.
	defer wg.Done()

	// Range over jobs channel; loop exits when jobs is closed and drained.
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(500 * time.Millisecond) // simulate work
		results <- job * 2                 // send result
	}
}

func main() {
	const numJobs = 10
	const numWorkers = 3

	// Buffered channels to reduce blocking.
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// Start the workers (fan-out).
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs to the workers.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // No more jobs; workers will exit after processing remaining jobs.

	// Wait for all workers to finish, then close results channel.
	// This must be done in a separate goroutine to avoid deadlock:
	// the main goroutine needs to continue to read results.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results (fan-in). The loop exits when results is closed.
	for res := range results {
		fmt.Println("Result:", res)
	}
}
