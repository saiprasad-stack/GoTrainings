// ============================== FILE 1 ==============================
// os_example.go
package main

import (
	"fmt"
	"os"
)

/*
os PACKAGE (File + OS operations)

Methods:
- Open
  Opens file in read-only mode

- Create
  Creates or truncates a file

- OpenFile
  Opens file with custom flags and permissions

- ReadFile
  Reads entire file into memory

- WriteFile
  Writes data to a file

- WriteString
  Writes string to file

- Close
  Closes the file descriptor
*/

func main() {
	file, _ := os.Create("example.txt")
	defer file.Close()

	file.WriteString("Hello OS\n")

	data, _ := os.ReadFile("example.txt")
	fmt.Println(string(data))
}
