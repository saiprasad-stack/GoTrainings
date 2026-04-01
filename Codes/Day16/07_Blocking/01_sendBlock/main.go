package main

import (
	"fmt"
	"time"
)

func main() {
	// Unbuffered channel – send blocks until receiver ready
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Receiver starting to receive...")
		<-ch
	}()

	fmt.Println("Sender about to block...")
	ch <- 999 // blocks for ~2 seconds
	fmt.Println("Sender unblocked")
}
