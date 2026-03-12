package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

type Operation func(int, int) int

func calculator(a int, b int, op Operation) int {
	return op(a, b)
}

func main() {
	fmt.Println("Higher Order functions")
	fmt.Println(add(23, 45))

	fmt.Println(calculator(23, 46, add))

}
