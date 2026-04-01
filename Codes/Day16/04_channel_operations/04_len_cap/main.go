package main

import "fmt"

func main() {
	ch := make(chan int)
	fmt.Printf("After make: len=%d\n", len(ch))

	go func() {
		fmt.Println("Received:", <-ch)
	}()

	ch <- 20

	fmt.Printf("After send: len=%d\n", len(ch))
}
