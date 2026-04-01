package main

import (
	"fmt"
	"os"
)

/*
fmt PACKAGE (Formatted I/O)

Methods:
- Fprintf
  Writes formatted output to a writer

- Fprintln
  Writes data with newline to a writer

- Fscan
  Reads formatted input from a reader

- Printf
  Prints formatted output to standard output
*/

func main() {
	file, _ := os.Create("formatted.txt")
	defer file.Close()

	fmt.Fprintf(file, "Name: %s Age: %d\n", "Go", 15)
}
