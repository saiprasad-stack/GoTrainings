package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch)

	// Receiving from closed channel yields zero value
	fmt.Println(<-ch) // 0

	// Double close panics
	// close(ch) // panic: close of closed channel

	// Sending on closed channel panics
	// ch <- 10 // panic: send on closed channel

	// Check if closed with comma-ok
	v, ok := <-ch
	fmt.Printf("v=%d, ok=%t\n", v, ok) // v=0, ok=false
}
