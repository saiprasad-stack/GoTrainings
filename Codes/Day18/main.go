package main

import (
	"fmt"
	"unicode/utf8"
)

/*
========================================
        STRINGS IN GO
========================================

1. Definition
- A string in Go is a read-only slice of bytes.
- It represents UTF-8 encoded text.

2. Key Properties
- Immutable (cannot be changed)
- UTF-8 encoded
- Indexing gives bytes, not characters

========================================
*/

func main() {

	/*
		========================================
		1. CREATING STRINGS
		========================================
	*/

	// Using double quotes
	str1 := "Hello"

	// Using var
	var str2 string = "World"

	// Using backticks (raw string)
	str3 := `This is
a multi-line
string`

	fmt.Println("str1:", str1)
	fmt.Println("str2:", str2)
	fmt.Println("str3:", str3)

	/*
		========================================
		2. STRING ASSIGNMENT
		========================================
	*/

	str := "Go"
	fmt.Println("Original:", str)

	str = "Golang" // reassignment
	fmt.Println("Updated:", str)

	/*
		========================================
		3. STRING LENGTH
		========================================
	*/

	fmt.Println("Length (bytes):", len("Hello"))
	fmt.Println("Length (bytes with unicode):", len("你好"))

	/*
		========================================
		4. INDEXING (BYTE LEVEL)
		========================================
	*/

	s := "Hello"
	fmt.Println("s[0]:", s[0]) // byte value
	fmt.Println("s[0] as char:", string(s[0]))

	/*
		========================================
		5. IMMUTABILITY
		========================================
	*/

	// s[0] = 'h' ❌ Not allowed (compile error)

	// To modify, convert to []byte
	b := []byte(s)
	b[0] = 'h'
	newStr := string(b)
	fmt.Println("Modified string:", newStr)

	/*
		========================================
		6. BYTE SLICE ([]byte)
		========================================
	*/

	strByte := "Hello"
	byteSlice := []byte(strByte)

	fmt.Println("Byte slice:", byteSlice)

	// Modify
	byteSlice[1] = 'a'
	fmt.Println("Modified byte slice:", string(byteSlice))

	/*
		========================================
		7. RUNE (UNICODE CODE POINT)
		========================================
	*/

	// Rune represents a Unicode character (int32)
	r := 'A'
	fmt.Printf("Rune: %c, Type: %T\n", r, r)

	/*
		========================================
		8. RUNE SLICE ([]rune)
		========================================
	*/

	unicodeStr := "Hello你好"

	runeSlice := []rune(unicodeStr)

	fmt.Println("Rune slice:", runeSlice)
	fmt.Println("Length (runes):", len(runeSlice))

	// Modify unicode safely
	runeSlice[0] = 'h'
	fmt.Println("Modified rune slice:", string(runeSlice))

	/*
		========================================
		9. ITERATING STRINGS
		========================================
	*/

	fmt.Println("\n--- Byte iteration ---")
	for i := 0; i < len(unicodeStr); i++ {
		fmt.Printf("%c ", unicodeStr[i]) // incorrect for unicode
	}

	fmt.Println("\n--- Rune iteration (correct) ---")
	for i, r := range unicodeStr {
		fmt.Printf("Index: %d, Rune: %c\n", i, r)
	}

	/*
		========================================
		10. UTF-8 FUNCTIONS
		========================================
	*/

	fmt.Println("Byte length:", len(unicodeStr))
	fmt.Println("Rune count:", utf8.RuneCountInString(unicodeStr))

	/*
		========================================
		11. STRING COMPARISON
		========================================
	*/

	a := "apple"
	bb := "banana"

	fmt.Println("a == bb:", a == bb)
	fmt.Println("a < bb:", a < bb) // lexicographical

	/*
		========================================
		12. STRING CONCATENATION
		========================================
	*/

	first := "Go"
	second := "Lang"

	result := first + second
	fmt.Println("Concatenated:", result)

	/*
		========================================
		13. STRING BUILDING (EFFICIENT WAY)
		========================================
	*/

	builder := make([]byte, 0)

	builder = append(builder, 'H')
	builder = append(builder, 'i')

	builtStr := string(builder)
	fmt.Println("Built string:", builtStr)

	/*
		========================================
		14. STRING TO BYTE / RUNE CONVERSION
		========================================
	*/

	text := "Hello"

	// string → []byte
	bytes := []byte(text)

	// string → []rune
	runes := []rune(text)

	fmt.Println("Bytes:", bytes)
	fmt.Println("Runes:", runes)

	/*
		========================================
		15. BYTE VS RUNE SUMMARY
		========================================

		- byte  = uint8 (1 byte)
		- rune  = int32 (Unicode character)
		- string = sequence of bytes (UTF-8)

		Use:
		- []byte → for performance / ASCII
		- []rune → for Unicode safety
	*/

}
