// ============================== FILE 3 ==============================
// bufio_example.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
bufio PACKAGE (Buffered I/O)

Methods:
- NewReader
  Creates a buffered reader

- ReadString
  Reads data until a delimiter

- NewWriter
  Creates a buffered writer

- WriteString
  Writes string using buffer

- Flush
  Flushes buffered data to destination

- NewScanner
  Creates scanner for token-based reading

- Scan
  Advances scanner to next token

- Text
  Returns current scanned token
*/

func main() {
	file, _ := os.Open("example.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
