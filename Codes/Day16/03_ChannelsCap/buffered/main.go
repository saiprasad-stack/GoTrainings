package main

import "fmt"

func main() {
	ch := make(chan int, 3) // buffered, capacity 3

	ch <- 1
	ch <- 2
	ch <- 3
	// ch <- 4 // would block because buffer is full

	fmt.Println("Buffer length:", len(ch))   // 3
	fmt.Println("Buffer capacity:", cap(ch)) // 3

	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
	fmt.Println(<-ch) // 3
}
