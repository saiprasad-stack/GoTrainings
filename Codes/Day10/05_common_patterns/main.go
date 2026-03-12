// ============================================================================
// Section 4: Common Patterns & Practical Operations
// ============================================================================
//
// THEORY:
// - Delete element at index i (preserving order):
//      s = append(s[:i], s[i+1:]...)
// - Delete without preserving order (swap with last):
//      s[i] = s[len(s)-1]; s = s[:len(s)-1]
// - Pop front: x, s := s[0], s[1:]
// - Pop back:  x, s := s[len(s)-1], s[:len(s)-1]
// - Insert at index i: s = append(s[:i], append([]T{x}, s[i:]...)...)
// - In-place filtering:
//      n := 0
//      for _, v := range s { if keep(v) { s[n] = v; n++ } }
//      s = s[:n]
// - Sorting: use sort.Ints, sort.Strings, or sort.Slice with custom less func.
// ============================================================================

package main

import (
	"fmt"
	"sort"
)

func main() {
	// Delete at index 1 (preserving order)
	s := []int{10, 20, 30, 40}
	i := 1
	s = append(s[:i], s[i+1:]...)
	fmt.Println("after delete index 1 (preserve order):", s) // [10 30 40]
	// Output: after delete index 1 (preserve order): [10 30 40]

	// Delete without order (swap with last)
	s = []int{10, 20, 30, 40}
	i = 1
	s[i] = s[len(s)-1]
	s = s[:len(s)-1]
	fmt.Println("after delete index 1 (unordered):", s) // [10 40 30] (order may vary)
	// Output: after delete index 1 (unordered): [10 40 30]

	// Pop front
	s = []int{5, 6, 7}
	first := s[0]
	s = s[1:]
	fmt.Printf("popped front %d, remaining %v\n", first, s)
	// Output: popped front 5, remaining [6 7]

	// Pop back
	s = []int{5, 6, 7}
	last := s[len(s)-1]
	s = s[:len(s)-1]
	fmt.Printf("popped back %d, remaining %v\n", last, s)
	// Output: popped back 7, remaining [5 6]

	// Insert at index 1
	s = []int{1, 2, 4}
	val := 3
	s = append(s[:1], append([]int{val}, s[1:]...)...)
	fmt.Println("after insert at index 1:", s) // [1 3 2 4]
	// Output: after insert at index 1: [1 3 2 4]

	// In-place filter (keep even numbers)
	nums := []int{1, 2, 3, 4, 5, 6}
	n := 0
	for _, v := range nums {
		if v%2 == 0 {
			nums[n] = v
			n++
		}
	}
	nums = nums[:n]
	fmt.Println("filtered evens:", nums) // [2 4 6]
	// Output: filtered evens: [2 4 6]

	// Sorting
	unsorted := []int{42, 13, 8, 99}
	sort.Ints(unsorted)
	fmt.Println("sorted ints:", unsorted)
	// Output: sorted ints: [8 13 42 99]

	strs := []string{"banana", "apple", "cherry"}
	sort.Strings(strs)
	fmt.Println("sorted strings:", strs)
	// Output: sorted strings: [apple banana cherry]

	// Custom sort (by length)
	words := []string{"go", "rust", "c", "python"}
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	fmt.Println("sorted by length:", words)
	// Output: sorted by length: [c go rust python]
}
