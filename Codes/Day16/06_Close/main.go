package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	// Receive from closed channel: buffered values come first, then zero values
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
	fmt.Println(<-ch) // 0 (channel closed, no blocking)

	// Check status with comma-ok
	v, ok := <-ch
	fmt.Printf("v=%d, ok=%t\n", v, ok) // v=0, ok=false
}
