package main

import "fmt"

/*
CONCATENATION

- Using +
- Avoid repeated concatenation in loops
*/

func main() {
	a := "Hello"
	b := "World"

	fmt.Println(a + " " + b)

	// Inefficient
	s := ""
	for i := 0; i < 5; i++ {
		s += "a"
	}

	fmt.Println("Loop concat:", s)
}
