package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
	close(ch)

	// Receive with status (comma-ok)
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel closed, no more values")
			break
		}
		fmt.Println("Received:", val, "ok =", ok)
	}
}
