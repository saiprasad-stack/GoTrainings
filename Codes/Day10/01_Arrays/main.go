// ============================================================================
// Arrays vs Slices in Go: Simple Comparison
// ============================================================================

package main

import "fmt"

func main() {
	// 1. Declaration: fixed size vs dynamic
	fmt.Println("=== 1. Declaration ===")
	var arr [3]int = [3]int{10, 20, 30} // array: size is part of type
	var slc []int = []int{10, 20, 30}   // slice: no size in type
	fmt.Printf("array: %v (type %T)\n", arr, arr)
	fmt.Printf("slice: %v (type %T)\n", slc, slc)
	// Output:
	// array: [10 20 30] (type [3]int)
	// slice: [10 20 30] (type []int)

	// 2. Type identity: size matters for arrays, not for slices
	fmt.Println("\n=== 2. Type Identity ===")
	var arr2 [3]int
	_ = arr2
	// arr2 = [4]int{1,2,3,4} // compile error: cannot use [4]int as [3]int
	fmt.Println("Arrays of different sizes are different types.")
	var slc2 []int
	slc2 = []int{1, 2, 3, 4} // fine, all []int are same type
	fmt.Printf("slice can be reassigned to different length: %v\n", slc2)

	// 3. Passing to functions: array copied, slice reference
	fmt.Println("\n=== 3. Passing to Functions ===")
	arr3 := [3]int{1, 2, 3}
	slc = []int{1, 2, 3}
	fmt.Println("Before function call:")
	fmt.Println("  array:", arr3)
	fmt.Println("  slice:", slc)

	modifyArray(arr3) // passes a copy
	modifySlice(slc)  // passes a reference (header copy, same underlying array)

	fmt.Println("After function call:")
	fmt.Println("  array (unchanged):", arr3)
	fmt.Println("  slice (changed):", slc)
	// Output:
	// Before function call:
	//   array: [1 2 3]
	//   slice: [1 2 3]
	// Inside modifyArray (copy): [99 2 3]
	// Inside modifySlice (reference): [99 2 3]
	// After function call:
	//   array (unchanged): [1 2 3]
	//   slice (changed): [99 2 3]

	// 4. Appending: only slices can grow
	fmt.Println("\n=== 4. Appending ===")
	// arr = append(arr, 4) // compile error: first argument to append must be slice
	slc = append(slc, 4, 5)
	fmt.Println("slice after append:", slc) // [99 2 3 4 5]
	// Output: slice after append: [99 2 3 4 5]

	// 5. Slicing: both can be sliced, result is always a slice
	fmt.Println("\n=== 5. Slicing ===")
	array := [5]int{1, 2, 3, 4, 5}
	sliceFromArray := array[1:4] // slice view into array
	fmt.Println("original array:", array)
	fmt.Println("slice from array:", sliceFromArray)

	// Modifying the slice modifies the original array
	sliceFromArray[0] = 99
	fmt.Println("array after modifying slice:", array)
	// Output:
	// original array: [1 2 3 4 5]
	// slice from array: [2 3 4]
	// array after modifying slice: [1 99 3 4 5]

	// 6. Comparability: arrays are comparable, slices are not
	fmt.Println("\n=== 6. Comparability ===")
	a1 := [2]int{1, 2}
	a2 := [2]int{1, 2}
	a3 := [2]int{3, 4}
	fmt.Println("array a1 == a2:", a1 == a2) // true
	fmt.Println("array a1 == a3:", a1 == a3) // false

	s1 := []int{1, 2}
	s2 := []int{1, 2}
	// fmt.Println(s1 == s2) // compile error: slice can only be compared to nil
	fmt.Println("slice s1 == nil:", s1 == nil) // false
	fmt.Println("slice s2 == nil:", s2 == nil) // false
	// Output:
	// array a1 == a2: true
	// array a1 == a3: false
	// slice s1 == nil: false
	// slice s2 == nil: false
}

// modifyArray receives a copy of the array
func modifyArray(arr [3]int) {
	arr[0] = 99
	fmt.Println("Inside modifyArray (copy):", arr)
}

// modifySlice receives a copy of the slice header, but the underlying array is shared
func modifySlice(slc []int) {
	slc[0] = 99
	fmt.Println("Inside modifySlice (reference):", slc)
}
