// goroutines

// Concurrency vs Parallelism --

package main

import (
	"fmt"
	"time"
)

func PrintNumbers1() {
	for i := 0; i <= 6; i++ {
		fmt.Println("Number 1 :", i)
	}
}
func PrintNumbers2() {
	for i := 0; i <= 6; i++ {
		fmt.Println("Number 2:", i)
	}
}

func main() {
	go PrintNumbers1()
	fmt.Println("Hello in GO ")
	go PrintNumbers2()
	time.Sleep(10 * time.Millisecond)

}

// goroutines -- A light weight thread
// by go keyword
// 2kb
// Why goroutines are fast ?
// random order of execution
