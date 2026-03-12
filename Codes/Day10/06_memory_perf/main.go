// ============================================================================
// Section 5: Performance & Memory Management
// ============================================================================
//
// THEORY:
// - Pre-allocation: using make([]T, 0, capacity) avoids repeated allocations.
// - Memory leaks: slicing a large array keeps the whole array alive.
//      * Use copy or full slice expressions to avoid this.
// - Growth policy: when append needs to allocate, capacity roughly doubles
//   for small sizes, then increases by about 25% for larger ones.
// - Zero-copy slicing: manipulating byte slices without copying improves perf.
// ============================================================================

package main

import "fmt"

func main() {
	// Pre-allocation demonstration (conceptual)
	const size = 1000
	// Without pre-allocation
	s1 := []int{}
	for i := 0; i < size; i++ {
		s1 = append(s1, i)
	}
	fmt.Printf("without pre-alloc: final cap=%d\n", cap(s1))
	// With pre-allocation
	s2 := make([]int, 0, size)
	for i := 0; i < size; i++ {
		s2 = append(s2, i)
	}
	fmt.Printf("with pre-alloc: final cap=%d\n", cap(s2))
	// Output (approximate):
	// without pre-alloc: final cap=1024
	// with pre-alloc: final cap=1000

	// Memory leak example
	largeArray := make([]int, 1e6) // 1 million ints
	for i := range largeArray {
		largeArray[i] = i
	}
	// Leak: small slice references large array
	leakySlice := largeArray[:10]
	fmt.Printf("leakySlice len=%d cap=%d (keeps largeArray alive)\n", len(leakySlice), cap(leakySlice))
	// Output: leakySlice len=10 cap=1000000 (keeps largeArray alive)

	// Fix: copy the needed part
	fixedSlice := make([]int, 10)
	copy(fixedSlice, largeArray[:10])
	fmt.Printf("fixedSlice len=%d cap=%d (independent)\n", len(fixedSlice), cap(fixedSlice))
	// Output: fixedSlice len=10 cap=10 (independent)

	// Growth policy observation
	fmt.Println("Capacity growth on append:")
	var growthSlice []int
	prevCap := 0
	for i := 0; i < 1000; i++ {
		growthSlice = append(growthSlice, i)
		if cap(growthSlice) != prevCap {
			fmt.Printf("  len=%-4d cap=%-4d\n", len(growthSlice), cap(growthSlice))
			prevCap = cap(growthSlice)
		}
	}
	// Output will show capacity jumps, e.g., 1,2,4,8,16,32,64,128,256,512,1024...
}
