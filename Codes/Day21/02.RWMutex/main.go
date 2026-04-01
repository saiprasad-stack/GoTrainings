package main

import (
	"fmt"
	"sync"
	"time"
)

/*
===========================
sync.RWMutex (Read-Write Mutex)
===========================

DEFINITION:
An RWMutex allows:
- MULTIPLE readers simultaneously
- ONLY ONE writer exclusively

WHY IMPORTANT:
- In many systems, reads >> writes
- RWMutex improves performance over Mutex in such cases

LOCK TYPES:
- mu.RLock()   → shared read lock
- mu.RUnlock() → release read lock
- mu.Lock()    → exclusive write lock
- mu.Unlock()  → release write lock

RULES:
1. Writers block ALL readers and writers
2. Readers block writers (until all readers release)
3. Do NOT upgrade RLock → Lock (can cause deadlock)

PITFALL:
- If writes are frequent → RWMutex may be slower than Mutex
*/

func main() {
	var mu sync.RWMutex
	data := 0

	var wg sync.WaitGroup

	// Writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		data = 100
		time.Sleep(time.Millisecond * 100)
		mu.Unlock()
	}()

	// Readers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu.RLock()
			fmt.Println("Reader", id, "reads:", data)
			mu.RUnlock()
		}(i)
	}

	wg.Wait()
}
