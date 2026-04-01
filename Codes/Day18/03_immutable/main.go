package main

import "fmt"

/*
IMMUTABILITY

- Strings cannot be modified directly
- Must convert to []byte or []rune
*/

func main() {
	s := "hello"

	// s[0] = 'H' // ❌ not allowed

	// Correct way
	b := []byte(s)
	b[0] = 'J'
	s = string(b)

	fmt.Println("Modified:", s)
}
