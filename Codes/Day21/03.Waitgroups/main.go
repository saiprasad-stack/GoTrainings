package main

import (
	"fmt"
	"sync"
)

/*
===========================
sync.WaitGroup
===========================

DEFINITION:
WaitGroup is used to wait for a collection of goroutines to finish.

HOW IT WORKS:
- Internal counter tracks number of goroutines

METHODS:
- wg.Add(n) → increment counter
- wg.Done() → decrement counter
- wg.Wait() → block until counter = 0

RULES:
1. ALWAYS call Add BEFORE starting goroutine
2. ALWAYS call Done (use defer wg.Done())
3. Do NOT reuse WaitGroup incorrectly (reset properly)

COMMON MISTAKE:
Calling Add inside goroutine → race condition

USE CASES:
- Parallel processing
- Batch jobs
- Worker coordination
*/

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			fmt.Println("Worker", id, "done")
		}(i)
	}

	wg.Wait()
	fmt.Println("All workers finished")
}
