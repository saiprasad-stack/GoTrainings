// ============================================================================
// Section 6: Advanced Topics & Edge Cases
// ============================================================================
//
// THEORY:
// - Full slice expression: s[low:high:max] sets capacity to max-low.
//      * Prevents accidental writes to original array.
// - Nil vs empty in JSON: nil slice marshals to null, empty to [].
// - Concurrency: slices are not safe for concurrent use without synchronization.
// - Reflection: reflect package can inspect/modify slices of unknown type.
// - Best practices: avoid pointers to slices, use copy for independence.
// ============================================================================

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
)

func main() {
	// Full slice expression
	originalArr := []int{1, 2, 3, 4, 5}
	fullSlice := originalArr[1:3:3] // length 2, capacity 2
	fmt.Printf("fullSlice=%v, len=%d, cap=%d\n", fullSlice, len(fullSlice), cap(fullSlice))
	// Output: fullSlice=[2 3], len=2, cap=2

	// Appending to fullSlice will allocate a new array because capacity is full
	fullSlice = append(fullSlice, 99)
	fmt.Println("originalArr after append to fullSlice:", originalArr) // unchanged
	// Output: originalArr after append to fullSlice: [1 2 3 4 5]

	// Nil vs Empty in JSON
	var nilSlice []int
	emptySlice := []int{}
	nilJSON, _ := json.Marshal(nilSlice)
	emptyJSON, _ := json.Marshal(emptySlice)
	fmt.Printf("nil slice JSON: %s\n", nilJSON)     // null
	fmt.Printf("empty slice JSON: %s\n", emptyJSON) // []
	// Output:
	// nil slice JSON: null
	// empty slice JSON: []

	// Concurrency example (with mutex)
	var mu sync.Mutex
	shared := make([]int, 0)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			mu.Lock()
			shared = append(shared, val)
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println("shared slice length (with mutex):", len(shared)) // 10
	// Output: shared slice length (with mutex): 10

	// Reflection: iterate over any slice
	anySlice := []string{"a", "b", "c"}
	v := reflect.ValueOf(anySlice)
	if v.Kind() == reflect.Slice {
		fmt.Println("Reflection iteration:")
		for i := 0; i < v.Len(); i++ {
			fmt.Printf("  Element %d: %v\n", i, v.Index(i).Interface())
		}
	}
	// Output:
	// Reflection iteration:
	//   Element 0: a
	//   Element 1: b
	//   Element 2: c

	// Best practice: avoid pointers to slices
	fmt.Println("Best practice: pass slices directly, not pointers to slices.")
	// Example: func process(s []int) { ... }   // correct
	//          func process(s *[]int) { ... }  // rarely needed
}
