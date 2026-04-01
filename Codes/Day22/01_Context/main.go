package main

import (
	"context"
	"fmt"
)

func main() {
	// context.Background() is the empty, non‑nil, immutable context.
	// It is used as the top‑level context for main(), tests, and at the start of incoming requests.
	bg := context.Background()
	fmt.Println("Background context:", bg)
	fmt.Println("Background context is empty:", bg.Value("anything") == nil) // true

	// context.TODO() is a placeholder when you’re unsure which context to use.
	// It signals that the code is not yet ready to accept a proper context.
	todo := context.TODO()
	fmt.Println("TODO context:", todo)

	// Both are identical in behavior: they have no deadline, cannot be cancelled,
	// and carry no values. They are never replaced; instead, you derive new contexts
	// from them using WithCancel, WithTimeout, WithDeadline, or WithValue.

	// Rule of thumb: Use Background() at the root of your program.
	//    Use TODO() temporarily while refactoring to add context support.
}
