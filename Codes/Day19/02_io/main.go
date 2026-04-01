package main

import (
	"fmt"
	"io"
	"os"
)

/*
io PACKAGE (Abstraction layer)

Methods:
- Copy
  Copies data from source to destination

- ReadAll
  Reads all data from a reader

- WriteString
  Writes string to a writer

Interfaces:
- Reader
  Defines read behavior for data sources

- Writer
  Defines write behavior for data destinations
*/

func main() {
	src, _ := os.Open("example.txt")
	defer src.Close()

	dst, _ := os.Create("copy.txt")
	defer dst.Close()

	bytes, _ := io.Copy(dst, src)
	fmt.Println("Copied:", bytes)
}
