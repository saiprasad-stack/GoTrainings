package main

import "fmt"

/*
STRING LITERALS

1. Interpreted (" ")
   - Supports escape sequences

2. Raw (` `)
   - No escaping
   - Preserves formatting
*/

func main() {
	// Interpreted
	s1 := "Hello\nWorld"
	fmt.Println(s1)

	// Raw
	s2 := `Hello\nWorld`
	fmt.Println(s2)
}
