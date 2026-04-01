package main

import "fmt"

func main() {
	ch := make(chan string, 3)
	ch <- "one"
	ch <- "two"
	ch <- "three"
	close(ch)

	// Range automatically stops after channel is closed and drained
	for msg := range ch {
		fmt.Println(msg)
	}
}
