package main

import (
	"fmt"
	"sync"
	"time"
)

/*
===========================
sync.Once
===========================

DEFINITION:
Ensures a function executes ONLY ONCE, even across multiple goroutines.

HOW IT WORKS:
- Internally uses atomic + mutex
- Guarantees thread-safe single execution

METHOD:
- once.Do(func())

USE CASES:
- Singleton pattern
- Lazy initialization
- Loading config, DB connection

IMPORTANT:
- If function panics → it is considered "done"
- Future calls will NOT retry

BEST PRACTICE:
Use with idempotent initialization logic
*/

func main() {
	var once sync.Once

	initFunc := func() {
		fmt.Println("Initialized ONLY ONCE")
	}

	for i := 0; i < 3; i++ {
		go once.Do(initFunc)
	}

	time.Sleep(time.Millisecond * 100)
}
