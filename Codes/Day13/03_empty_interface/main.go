package main

import "fmt"

func main() {
	// 1. Store different types in a slice of interface{}
	var things []interface{}
	things = append(things, 42)
	things = append(things, "hello")
	things = append(things, 3.14)
	things = append(things, true)

	// 2. Use type switch to handle each value
	fmt.Println("Type switch example:")
	for _, v := range things {
		switch val := v.(type) {
		case int:
			fmt.Printf("Integer: %d\n", val)
		case string:
			fmt.Printf("String: %s\n", val)
		case float64:
			fmt.Printf("Float: %f\n", val)
		case bool:
			fmt.Printf("Boolean: %v\n", val)
		default:
			fmt.Printf("Unknown type: %T\n", val)
		}
	}

	// 3. Type assertion with "comma ok"
	fmt.Println("\nType assertion example:")
	var x interface{} = "hello"
	str, ok := x.(string)
	if ok {
		fmt.Println("x is a string:", str)
	}
}
