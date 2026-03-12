// ============================================================================
// Creation & Basic Operations
// ============================================================================
//
// - Ways to create slices:
//      * Literal: s := []T{vals...}
//      * make: make([]T, length, capacity)  (capacity optional)
//      * From array: array[low:high]
//      * From another slice: slice[low:high]
// - Access elements with s[index] (0 <= index < len(s))
// - Slicing expressions:
//      * Simple: s[low:high] → length = high-low, capacity = cap(s)-low
//      * Full (with capacity control): s[low:high:max] → capacity = max-low
// - Iteration: classic for loop or for range (returns index and copy of element)
// - Passing to functions: slice header is passed by value (copy), but the
//   underlying array is shared. Modifying elements affects caller; appending
//   may not unless you return the new slice.
// ============================================================================

package main

import "fmt"

func main() {
	// Different creation methods
	a := []int{1, 2, 3}    // literal
	b := make([]int, 3, 5) // length 3, capacity 5
	arr := [4]int{10, 20, 30, 40}
	c := arr[1:3] // from array
	d := c[0:1]   // from slice

	fmt.Println("a:", a, "len:", len(a), "cap:", cap(a))
	fmt.Println("b:", b, "len:", len(b), "cap:", cap(b))
	fmt.Println("c:", c, "len:", len(c), "cap:", cap(c))
	fmt.Println("d:", d, "len:", len(d), "cap:", cap(d))
	// Output:
	// a: [1 2 3] len: 3 cap: 3
	// b: [0 0 0] len: 3 cap: 5
	// c: [20 30] len: 2 cap: 3
	// d: [20] len: 1 cap: 3

	// Access and modify (shared underlying array)
	c[0] = 99
	fmt.Println("arr after c[0] change:", arr) // underlying array changed
	// Output: arr after c[0] change: [10 99 30 40]

	// Iteration
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("Iterating over fruits:")
	for i, v := range fruits {
		fmt.Printf("  index %d: %s\n", i, v)
	}
	// Output:
	// Iterating over fruits:
	//   index 0: apple
	//   index 1: banana
	//   index 2: cherry

	// Function argument demonstration
	fmt.Println("Before modifySlice:", fruits)
	modifySlice(fruits)
	fmt.Println("After modifySlice:", fruits) // first element changed
	// Output:
	// Before modifySlice: [apple banana cherry]
	// Inside modifySlice: [KIWI banana cherry date]
	// After modifySlice: [KIWI banana cherry]
}

// modifySlice demonstrates slice passing semantics
func modifySlice(s []string) {
	if len(s) > 0 {
		s[0] = "KIWI" // modifies underlying array
	}
	s = append(s, "date") // this append only affects local copy
	fmt.Println("Inside modifySlice:", s)
}
