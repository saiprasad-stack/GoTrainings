/*
PROJECT FLOW:
- os → open files
- bufio → count lines
- io → copy data
- fmt → write summary
- filepath → manage paths
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	input := filepath.Join("data", "input.txt")
	output := filepath.Join("data", "output.txt")

	// Ensure directory exists
	err := os.MkdirAll("data", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// Open input file
	in, err := os.Open(input)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer in.Close()

	// Create output file
	out, err := os.Create(output)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer out.Close()

	// Count lines
	scanner := bufio.NewScanner(in)
	count := 0
	for scanner.Scan() {
		count++
	}

	// Reset file pointer
	_, err = in.Seek(0, 0)
	if err != nil {
		fmt.Println("Seek error:", err)
		return
	}

	// Copy data
	bytes, err := io.Copy(out, in)
	if err != nil {
		fmt.Println("Copy error:", err)
		return
	}

	// Write summary to file
	fmt.Fprintf(out, "\nLines: %d\nBytes: %d\n", count, bytes)

	// ALSO print to console (important)
	fmt.Println("Processing complete")
	fmt.Println("Lines:", count)
	fmt.Println("Bytes:", bytes)
}
