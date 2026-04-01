package main

import "fmt"

func main() {
	ch := make(chan int)

	select {
	case v := <-ch:
		fmt.Println("Received:", v)
	default:
		fmt.Println("No value available, non-blocking")
	}
}
