package main

import "fmt"

func main() {
	ch := make(chan string, 1)

	ch <- "ping"
	msg := <-ch

	fmt.Println(msg)
}
