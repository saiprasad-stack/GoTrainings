package main

import "fmt"

// ==========================================================
// DEFINITIONS
// ==========================================================
// Type Assertion: Used to extract the underlying concrete value
// from an interface variable. It asserts that the interface holds
// a specific type. Syntax: value.(Type) or value.(Type) with "comma ok".
//
// Type Switch: A switch statement that branches based on the type
// of an interface variable. It allows handling multiple possible types
// cleanly. Syntax: switch v := value.(type) { case Type1: ... }
// ==========================================================

func main() {
	fmt.Println("=== TYPE ASSERTION vs TYPE SWITCH ===")

	// Create an interface variable that holds different types
	var data interface{}

	// ==========================================================
	// PART 1: TYPE ASSERTION
	// ==========================================================
	fmt.Println("1. TYPE ASSERTION")

	data = "hello, world"

	// Unsafe assertion: if data is not a string, this panics!
	// s := data.(string) // (commented to avoid panic)

	// Safe assertion with "comma ok" (recommended)
	if str, ok := data.(string); ok {
		fmt.Printf("   Safe assertion: data is a string → %q\n", str)
	} else {
		fmt.Println("   data is not a string")
	}

	// Asserting the wrong type – safe version
	if num, ok := data.(int); ok {
		fmt.Printf("   data is an int: %d\n", num)
	} else {
		fmt.Println("   Safe assertion: data is NOT an int (as expected)")
	}

	// Another example with int
	data = 42
	if num, ok := data.(int); ok {
		fmt.Printf("   Safe assertion: data is an int → %d\n", num)
	}

	fmt.Println()

	// ==========================================================
	// PART 2: TYPE SWITCH
	// ==========================================================
	fmt.Println("2. TYPE SWITCH")

	// Create a slice of interface{} with mixed types
	mixed := []interface{}{
		"golang",
		100,
		3.14,
		true,
		[]int{1, 2, 3},
	}

	for idx, val := range mixed {
		fmt.Printf("   Item %d: ", idx)
		switch v := val.(type) { // special .(type) syntax only in switch
		case string:
			fmt.Printf("string → %q (length %d)\n", v, len(v))
		case int:
			fmt.Printf("int → %d\n", v)
		case float64:
			fmt.Printf("float64 → %f\n", v)
		case bool:
			fmt.Printf("bool → %v\n", v)
		default:
			fmt.Printf("unknown type %T → %v\n", v, v)
		}
	}

	fmt.Println()
	/*
	   Difference between Type Assertion and Type Switch

	   | Feature            | Type Assertion               | Type Switch                     |
	   |--------------------|------------------------------|---------------------------------|
	   | Purpose            | Extract single concrete type | Handle multiple possible types  |
	   | Syntax             | value.(Type)                 | switch value.(type) { case T: } |
	   | Number of cases    | One                          | Multiple                        |
	   | Panic risk         | Yes (if unchecked)           | No (safe by design)             |
	   | Use case           | You know exactly what type   | Type is unknown/variable        |
	   |                    | you expect                   |                                 |
	*/
}
