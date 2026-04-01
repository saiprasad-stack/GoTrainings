package main

import "fmt"

/*
ITERATION

1. Byte-wise (may break Unicode)
2. Rune-wise (safe)
*/

func main() {
	s := "Go语言"

	// Byte-wise
	for i := 0; i < len(s); i++ {
		fmt.Println("Byte:", s[i])
	}

	// Rune-wise
	for _, r := range s {
		fmt.Println("Rune:", string(r))
	}
}
