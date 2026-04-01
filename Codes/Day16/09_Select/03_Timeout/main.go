package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	select {
	case v := <-ch:
		fmt.Println("Received:", v)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout after 2 seconds")
	}
}
