package main

import "fmt"

/*
INDEXING

- s[i] gives byte, not character
*/

func main() {
	s := "hêllo"

	fmt.Println("Byte:", s[1])                // 104
	fmt.Println("Character:", string(s[1:3])) // h
}
