package main

import "fmt"

/*
BASICS OF STRINGS

- A string is a read-only slice of bytes
- Stored in UTF-8 encoding
- Immutable (cannot change after creation)
*/

func main() {
	s := "Hello"
	s2 := "Héllo"
	fmt.Println("String:", s2)
	fmt.Println("Length (bytes):", len(s2))
	fmt.Println("String:", s)
	fmt.Println("Length (bytes):", len(s))
}

/*
Unicode
A standard that defines characters
Gives each character a unique number (code point)

Example:
'A' → U+0041, 'ह' → U+0939

UTF-8
An encoding of Unicode
Stores characters as 1 to 4 bytes
Most widely used

UTF-16
Uses 2 or 4 bytes
Used in systems like Java, Windows

UTF-32
Fixed 4 bytes per character
Simple but memory heavy

ASCII
Old encoding
Only 128 characters
1 byte per character

*/
