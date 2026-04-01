package main

import (
	"fmt"
	"sync"
	"time"
)

// worker listens on the stop channel. When stop is closed, it exits.
func worker(id int, stop <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stop:
			// stop channel closed – time to exit.
			fmt.Printf("Worker %d stopping\n", id)
			return
		default:
			// do normal work
			fmt.Printf("Worker %d working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	// stop channel is used to broadcast a shutdown signal.
	stop := make(chan struct{})
	var wg sync.WaitGroup

	// Start several workers.
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, stop, &wg)
	}

	// Let them work for a while.
	time.Sleep(2 * time.Second)

	// Broadcast stop signal by closing the channel.
	// All workers receive the zero value and exit their select.
	fmt.Println("Broadcasting stop signal")
	close(stop)

	// Wait for all workers to finish.
	wg.Wait()
	fmt.Println("All workers stopped")
}
