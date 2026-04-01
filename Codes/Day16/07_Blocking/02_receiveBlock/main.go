package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Sender sending...")
		ch <- 123
	}()

	fmt.Println("Receiver about to block...")
	val := <-ch // blocks until value arrives
	fmt.Println("Received:", val)
}
