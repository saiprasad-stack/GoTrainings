package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 1. Atoi (string → int)
	// Converts string to int (base 10)
	i, _ := strconv.Atoi("123")
	fmt.Println(i) // 123

	i2, _ := strconv.Atoi("-45")
	fmt.Println(i2) // -45

	i3, err := strconv.Atoi("abc")
	fmt.Println(i3, err) // 0 error

	// 2. Itoa (int → string)
	// Converts int to string
	s := strconv.Itoa(123)
	fmt.Println(s) // "123"

	fmt.Println(strconv.Itoa(-10)) // "-10"
	fmt.Println(strconv.Itoa(0))   // "0"

	// 3. ParseInt
	// Converts string to integer with base and bit size
	n, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(n) // 123
	next, _ := strconv.ParseInt("123$@165", 10, 64)
	fmt.Println(next) // 123

	n2, _ := strconv.ParseInt("111", 2, 64) // binary
	fmt.Println(n2)                         // 7

	n3, err := strconv.ParseInt("xyz", 10, 64)
	fmt.Println(n3, err) // 0 error

	// 4. ParseFloat
	// Converts string to float
	f, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println(f) // 3.14

	f2, _ := strconv.ParseFloat("10", 64)
	fmt.Println(f2) // 10

	f3, err := strconv.ParseFloat("abc", 64)
	fmt.Println(f3, err) // 0 error

	// 5. ParseBool
	// Converts string to boolean
	b, _ := strconv.ParseBool("true")
	fmt.Println(b) // true

	b2, _ := strconv.ParseBool("false")
	fmt.Println(b2) // false

	b3, err := strconv.ParseBool("notbool")
	fmt.Println(b3, err) // false error

	// 6. FormatInt
	// Converts integer to string with given base
	fmt.Println(strconv.FormatInt(123, 10)) // "123"
	fmt.Println(strconv.FormatInt(10, 2))   // "1010" (binary)
	fmt.Println(strconv.FormatInt(255, 16)) // "ff" (hex)

	// 7. FormatFloat
	// Converts float to string
	// 'f' = decimal format, 2 = precision
	fmt.Println(strconv.FormatFloat(3.14, 'f', 2, 64))    // "3.14"
	fmt.Println(strconv.FormatFloat(3.14159, 'f', 3, 64)) // "3.142"
	fmt.Println(strconv.FormatFloat(10, 'f', 0, 64))      // "10"

	// 8. FormatBool
	// Converts bool to string
	fmt.Println(strconv.FormatBool(true))  // "true"
	fmt.Println(strconv.FormatBool(false)) // "false"
}
