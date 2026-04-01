package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int // nil channel

	select {
	case v := <-ch: // this case is never selected because ch is nil
		fmt.Println("Received:", v)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout (nil channel never ready)")
	}
}
