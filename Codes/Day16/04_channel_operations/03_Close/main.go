package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	close(ch)

	// Receive remaining values
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
	// Channel is now empty and closed
	fmt.Println(<-ch) // 0 (zero value) and does not block

	// Sending on a closed channel panics
	// ch <- 3 // panic: send on closed channel
}
