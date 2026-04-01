// 1. Use `sync.WaitGroup` to force `main` to wait.
// 2. Add(1) for each task. Call Done() when a task finishes.
//    Call Wait() in main to pause until the counter reaches zero.
// =====================================================================

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1) // Tell WaitGroup we have 1 pending task

	go func() {
		defer wg.Done() // Signal that this task is done when the function exits

		fmt.Println("Goroutine: Starting work...")
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine: Work complete!")
	}()

	fmt.Println("Main: Waiting for the goroutine to finish...")
	wg.Wait() // Pauses execution here until the counter is 0

	fmt.Println("Main: Goroutine finished, program exiting cleanly.")
}
