// ============================================================================
// Slice Fundamentals & Memory Model
// ============================================================================
//
// - A slice is a descriptor of a contiguous segment of an underlying array.
// - It provides a dynamic, flexible view into an array.
// - Internally, a slice is a small struct containing:
//      * Pointer: memory address of the first element
//      * Length: number of elements currently in the slice
//      * Capacity: maximum number of elements the slice can hold
//                (from the pointer to the end of the array)
// - Nil slice: var s []int → no underlying array, s == nil
// - Empty slice: s := []int{} or make([]int, 0) → has underlying array,
//                but length 0, s != nil
// ============================================================================

package main

import "fmt"

func main() {
	// Nil slice
	var s1 []int
	fmt.Println("s1:", s1, "len:", len(s1), "cap:", cap(s1), "is nil?", s1 == nil)
	// Output: s1: [] len: 0 cap: 0 is nil? true

	// Empty slice (literal)
	s2 := []int{}
	fmt.Println("s2:", s2, "len:", len(s2), "cap:", cap(s2), "is nil?", s2 == nil)
	// Output: s2: [] len: 0 cap: 0 is nil? false

	// Empty slice using make
	s3 := make([]int, 0)
	fmt.Println("s3:", s3, "len:", len(s3), "cap:", cap(s3), "is nil?", s3 == nil)
	// Output: s3: [] len: 0 cap: 0 is nil? false

	// Slice with elements
	s4 := []int{10, 20, 30}
	fmt.Println("s4:", s4, "len:", len(s4), "cap:", cap(s4))
	// Output: s4: [10 20 30] len: 3 cap: 3

	// Underlying array demonstration
	arr := [5]int{1, 2, 3, 4, 5}
	s5 := arr[1:4] // slice from index 1 to 3 (length 3)
	fmt.Printf("s5: %v, len=%d, cap=%d\n", s5, len(s5), cap(s5))
	// Output: s5: [2 3 4], len=3, cap=4  (capacity from index 1 to end of array = 4)
}
