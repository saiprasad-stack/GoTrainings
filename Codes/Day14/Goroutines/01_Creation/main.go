// 1. Creation: Use the `go` keyword before a function call.
// 2. The Rule: If `main` ends, all running goroutines are immediately
//  terminated (killed abruptly).
// =====================================================================

package main

import (
	"fmt"
	"time"
)

func Hello() {
	fmt.Println("Hello World")
}
func main() {
	// go func() {
	// 	// This will NEVER print because main finishes in 1 second.
	// 	fmt.Println("Goroutine: I finished my work!")
	// }()

	Hello()
	fmt.Println("Main: Doing some quick work...")

	time.Sleep(5 * time.Second)
	fmt.Println("Main: I'm done. Exiting now and killing the goroutine!")
}
