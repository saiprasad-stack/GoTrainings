package main

import "fmt"

// Employee struct definition
type Employee struct {
	id   int
	name string
}

// Value Receiver Method
func (e Employee) myFunc() {
	e.name = "Pal"
	fmt.Println("Inside myFunc (Value Receiver):", e.name)
}

// Pointer Receiver Method
func (e *Employee) myFuncPointer() {
	e.name = "Pal"
	fmt.Println("Inside myFuncPointer (Pointer Receiver):", e.name)
}

func main() {

	emp := Employee{123, "Arvinder"}
	emp2 := Employee{123, "Arvinder"}
	emp3 := Employee{123, "Pal"}
	emp4 := Employee{124, "Arvinder"}

	fmt.Println("Struct Comparison:")
	fmt.Println(emp == emp2)  // true
	fmt.Println(emp == emp3)  // false
	fmt.Println(emp == emp4)  // false
	fmt.Println(emp3 == emp4) // false

	fmt.Println()

	fmt.Println("Before calling myFunc:", emp.name)
	emp.myFunc()
	fmt.Println("After calling myFunc:", emp.name)

	fmt.Println()

	fmt.Println("Before calling myFuncPointer:", emp.name)
	emp.myFuncPointer()
	fmt.Println("After calling myFuncPointer:", emp.name)
}
