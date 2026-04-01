// 1. Blocking: Goroutines can block (pause) due to time.Sleep, I/O, or channels.
// 2. Efficiency: When blocked, the Go scheduler puts the goroutine to sleep
//    and immediately uses the underlying OS thread to run OTHER goroutines.
// =====================================================================

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine 1 (Simulates heavy blocking)
	go func() {
		defer wg.Done()
		fmt.Println("G1: Going to sleep for 2 seconds (Blocking)...")
		time.Sleep(2 * time.Second)
		fmt.Println("G1: Woke up!")
	}()

	// Goroutine 2 (Runs while G1 is blocked)
	go func() {
		defer wg.Done()
		fmt.Println("G2: G1 is blocked, so the scheduler is letting me run!")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("G2: I'm done early.")
	}()

	wg.Wait()
	fmt.Println("Main: All done.")
}
