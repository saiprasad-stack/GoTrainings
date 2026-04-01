package main

import (
	"fmt"
	"sync"
)

/*
===========================
sync.Mutex (Mutual Exclusion)
===========================

DEFINITION:
A Mutex is a locking mechanism used to ensure that only ONE goroutine
can access a shared resource (critical section) at a time.

WHY NEEDED:
- Go allows multiple goroutines to run concurrently
- If multiple goroutines modify shared data → RACE CONDITION
- Mutex prevents inconsistent/corrupted data

KEY METHODS:
- mu.Lock()   → acquire lock (blocks if already locked)
- mu.Unlock() → release lock

IMPORTANT RULES:
1. Always pair Lock() with Unlock()
2. Use defer mu.Unlock() for safety (recommended)
3. Do NOT copy a Mutex (pass by reference if needed)
4. Keep critical section SMALL (performance)

WHEN TO USE:
- Shared counters
- Shared maps/slices
- Any mutable shared state

ALTERNATIVE:
- Channels (idiomatic Go)
*/

func main() {
	var mu sync.Mutex
	counter := 0

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mu.Lock() // entering critical section
			counter++
			mu.Unlock() // leaving critical section
		}()
	}

	wg.Wait()
	fmt.Println("Mutex Counter:", counter)
}
