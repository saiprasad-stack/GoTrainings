package main

import "fmt"

/*
========================================
        BYTE vs RUNE IN GO
========================================

1. STRING INTERNALS
-------------------
- A string in Go is a sequence of BYTES
- These bytes are encoded using UTF-8

Example:
"Aé😊" → stored as bytes internally

----------------------------------------

2. WHAT IS A BYTE?
-------------------
- byte = uint8 (8 bits)
- Represents raw data
- Does NOT represent a full character (always)

Example:
'é' is NOT 1 byte → it is 2 bytes in UTF-8

----------------------------------------

3. WHAT IS A RUNE?
-------------------
- rune = int32
- Represents a Unicode code point (a character)

Example:
'é' → one rune

----------------------------------------

4. KEY DIFFERENCE
-------------------
byte → raw storage unit
rune → actual character

----------------------------------------

5. UTF-8 BEHAVIOR
-------------------
Character size varies:

'A'   → 1 byte
'é'   → 2 bytes
'ह'   → 3 bytes
'😊'  → 4 bytes

----------------------------------------

6. IMPORTANT RULES
-------------------
- len(string) → number of BYTES
- len([]rune(string)) → number of CHARACTERS

----------------------------------------

7. WHEN TO USE WHAT?
-------------------
Use byte:
- low-level operations
- networking, file reading

Use rune:
- text processing
- Unicode-safe operations

========================================
*/

func main() {

	/*
		========================================
		        EXAMPLE STRING
		========================================
	*/
	s := "Aé😊"

	fmt.Println("Original String:", s)

	/*
		========================================
		        BYTE LEVEL VIEW
		========================================
	*/
	fmt.Println("\n--- BYTE VIEW ---")

	fmt.Println("Total Bytes:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("Byte %d: %v\n", i, s[i])
	}

	/*
		Observation:
		- You will see multiple bytes for some characters
		- These are raw UTF-8 encoded values
	*/

	/*
		========================================
		        RUNE LEVEL VIEW
		========================================
	*/
	fmt.Println("\n--- RUNE VIEW ---")

	r := []rune(s)

	fmt.Println("Total Characters (Runes):", len(r))

	for i, v := range r {
		fmt.Printf("Rune %d: %c (Unicode: %U)\n", i, v, v)
	}

	/*
		Observation:
		- Each rune represents one complete character
		- Unicode code point is shown using %U
	*/

	/*
		========================================
		        LENGTH DIFFERENCE
		========================================
	*/
	fmt.Println("\n--- LENGTH COMPARISON ---")

	fmt.Println("len(s) (bytes):", len(s))
	fmt.Println("len([]rune(s)) (characters):", len([]rune(s)))

	/*
		========================================
		        INDEXING PITFALL
		========================================
	*/
	fmt.Println("\n--- INDEXING PITFALL ---")

	s2 := "😊"

	fmt.Println("Bytes in 😊:", len(s2))

	// Wrong way
	fmt.Println("First byte (wrong):", s2[0])
	fmt.Println("As string (wrong):", string(s2[0]))

	// Correct way
	r2 := []rune(s2)
	fmt.Println("Correct character:", string(r2[0]))

	/*
		========================================
		        ITERATION COMPARISON
		========================================
	*/
	fmt.Println("\n--- ITERATION COMPARISON ---")

	str := "Go语言"

	fmt.Println("\nByte-wise iteration:")
	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}

	fmt.Println("\nRune-wise iteration:")
	for _, r := range str {
		fmt.Println(string(r))
	}

	/*
		========================================
		        FINAL TAKEAWAYS
		========================================

		1. Strings are stored as UTF-8 bytes
		2. byte = raw data (may be part of a character)
		3. rune = actual character
		4. len() gives bytes, not characters
		5. Use rune for Unicode-safe operations
		6. Indexing a string gives a byte, not a character

		========================================
	*/
}
