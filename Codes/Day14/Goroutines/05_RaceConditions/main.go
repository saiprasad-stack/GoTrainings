// =====================================================================
// GOROUTINE EXECUTION ORDER: QUICK REFERENCE
// =====================================================================
// THEORY:
// 1. Order? NO. Goroutines execute in a completely random, unpredictable order.
// 2. Why? The Go scheduler runs them non-deterministically based on CPU availability.
// 3. The Rule: Never assume Task A will finish before Task B just because you
//    launched Task A first.
// 4. The Fix: If you MUST have a specific order, you have to force it yourself
//    using synchronization tools like Channels.
// =====================================================================

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// --- 1. THE PROBLEM: RANDOM EXECUTION ORDER ---
	// If you run this block multiple times, the output order will constantly change.
	fmt.Println("--- 1. Random Order (Run multiple times to see) ---")

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Unordered Task %d finished\n", id) // Who finishes first? Nobody knows.
		}(i)
	}
	wg.Wait()

	// --- 2. THE FIX: ENFORCING ORDER WITH CHANNELS ---
	// We use a channel to make Task B wait patiently until Task A says it is okay to go.
	fmt.Println("\n--- 2. Enforcing Order ---")

	signalChan := make(chan bool) // A channel acting as a simple signal relay

	// Task B (We launch it first, but we force it to wait)
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-signalChan // BLOCKS here until a signal comes through the channel
		fmt.Println("Task B: I am running ONLY because Task A gave me the signal.")
	}()

	// Task A (We launch it second, but it does its work first)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Task A: Doing my work first...")
		signalChan <- true // Sends a signal into the channel, unblocking Task B
	}()

	wg.Wait()
	fmt.Println("Main: All done.")
}
