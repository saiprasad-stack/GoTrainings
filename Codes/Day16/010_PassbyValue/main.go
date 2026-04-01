package main

import "fmt"

type Data struct {
	ID int
}

func main() {
	ch := make(chan Data, 1)

	original := Data{ID: 42}
	ch <- original // a copy is sent

	// Modify the original after sending
	original.ID = 100

	received := <-ch
	fmt.Println("Received ID:", received.ID) // still 42, not 100
}
