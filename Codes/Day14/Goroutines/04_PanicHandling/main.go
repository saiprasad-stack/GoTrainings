// 1. A `panic` inside a goroutine crashes the ENTIRE program.
// 2. You CANNOT catch a goroutine's panic from `main`.
// 3. You must use `recover()` inside the failing goroutine itself.

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		// The recovery mechanism MUST be deferred inside the goroutine
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Goroutine: Caught a panic ->", r)
			}
		}()

		fmt.Println("Goroutine: About to do something dangerous...")
		panic("CRITICAL FAILURE") // Triggers the panic

		// fmt.Println("This line will never run.")
	}()

	wg.Wait()
	fmt.Println("Main: Program survived the panic and exited normally.")
}
