package main

import "fmt"

/*
MODIFYING STRINGS SAFELY

- Convert to []rune for Unicode safety
*/

func main() {
	s := "hello 世界"

	r := []rune(s)
	r[0] = 'H'

	s = string(r)

	fmt.Print("Modified:", s[6:])
}
