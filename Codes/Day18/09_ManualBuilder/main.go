package main

import "fmt"

/*
========================================
    MANUAL STRING BUILD (SIMPLE WAY)
========================================

Goal:
- Build a string step by step using bytes

Why?
- Strings are immutable (cannot be changed)
- Using + in loops creates many temporary strings (slow)
- Using []byte is more efficient

----------------------------------------

CORE IDEA:

1. Create empty byte slice
2. Add characters (as bytes)
3. Convert to string

----------------------------------------

Think like this:

[]byte → putting letters in a box
string() → turning that box into a word

========================================
*/

func main() {

	// STEP 1: Create empty container
	builder := make([]byte, 0)

	/*
		At this point:
		builder = []
		(empty list of bytes)
	*/

	// STEP 2: Add characters one by one
	builder = append(builder, 'H') // 'H' → byte (72)
	builder = append(builder, 'i') // 'i' → byte (105)

	/*
		Now internally:
		builder = [72 105]

		These are ASCII values:
		H = 72
		i = 105
	*/

	// See raw bytes
	fmt.Println("Raw bytes:", builder)

	// STEP 3: Convert bytes → string
	s := string(builder)

	/*
		Now:
		[72 105] → "Hi"
	*/

	// Final output
	fmt.Println("Built string:", s)

	/*
		========================================
		        IMPORTANT NOTES
		========================================

		1. 'H' vs "H"
		   - 'H' → character (byte)
		   - "H" → string

		2. []byte is NOT a string
		   - Must convert using string()

		3. This method is useful when:
		   - Building strings in loops
		   - Performance matters

		========================================
	*/

	// Example: building in loop
	fmt.Println("\nBuilding in loop:")

	temp := make([]byte, 0)

	for i := 0; i < 5; i++ {
		temp = append(temp, 'a')
	}

	fmt.Println("Loop result:", string(temp))
}
