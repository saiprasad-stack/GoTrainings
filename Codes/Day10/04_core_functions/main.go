// ============================================================================
// Core Manipulation Functions
// ============================================================================
//
// - append(s, elements...) : appends elements to slice.
//      * If capacity is sufficient, it adds to existing underlying array.
//      * If capacity is insufficient, it allocates a new larger array and copies.
// - copy(dst, src) : copies elements from src to dst, returns number copied.
//      * It copies the minimum of len(dst) and len(src).
// - clear(s) (Go 1.21+) : sets all elements up to length to zero values.
//      * Does not change length or capacity.
// ============================================================================

package main

import "fmt"

func main() {
	// Append within capacity
	s := make([]int, 2, 4) // len=2, cap=4
	s[0] = 1
	s[1] = 2
	fmt.Println("before append:", s, "len:", len(s), "cap:", cap(s))
	// Output: before append: [1 2] len: 2 cap: 4

	s = append(s, 3, 4) // within capacity, no new array
	fmt.Println("after append 3,4:", s, "len:", len(s), "cap:", cap(s))
	// Output: after append 3,4: [1 2 3 4] len: 4 cap: 4

	s = append(s, 5) // exceeds capacity, new array allocated
	fmt.Println("after append 5:", s, "len:", len(s), "cap:", cap(s))
	// Output: after append 5: [1 2 3 4 5] len: 5 cap: 8 (typical growth)

	// Append multiple and slice concatenation
	s2 := []int{10, 20}
	s = append(s, s2...) // note the ...
	fmt.Println("after appending s2:", s)
	// Output: after appending s2: [1 2 3 4 5 10 20]

	// Copy
	src := []int{100, 200, 300}
	dst := make([]int, 2) // length 2
	n := copy(dst, src)
	fmt.Printf("copied %d elements: dst = %v\n", n, dst)
	// Output: copied 2 elements: dst = [100 200]

	// Full copy (deep copy)
	original := []int{1, 2, 3}
	copied := make([]int, len(original))
	copy(copied, original)
	copied[0] = 999
	fmt.Println("original:", original, "copied:", copied)
	// Output: original: [1 2 3] copied: [999 2 3]

	// Clear (manual equivalent; in Go 1.21+ you could use clear(s))
	for i := range s {
		s[i] = 0
	}
	fmt.Println("after manual clear:", s)
	// Output: after manual clear: [0 0 0 0 0 0 0]
}
