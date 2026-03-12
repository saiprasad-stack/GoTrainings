package main

import (
	"fmt"
	"strings"
)

// Reader defines read behavior
type Reader interface {
	Read() (string, error)
}

// Writer defines write behavior
type Writer interface {
	Write(data string) error
}

// ReadWriter composes Reader and Writer
type ReadWriter interface {
	Reader
	Writer
	// Greet() string
}

// Closer defines close behavior
type Closer interface {
	Close() error
}

// ReadWriteCloser composes multiple interfaces
type ReadWriteCloser interface {
	ReadWriter
	Closer
}

// File represents a simple file (like in os package)
type File struct {
	Name    string
	content []string
	closed  bool
}

func NewFile(name string) *File {
	return &File{
		Name:    name,
		content: make([]string, 0),
	}
}

func (f *File) Read() (string, error) {
	if f.closed {
		return "", fmt.Errorf("file is closed")
	}
	return strings.Join(f.content, "\n"), nil
}

func (f *File) Write(data string) error {
	if f.closed {
		return fmt.Errorf("file is closed")
	}
	f.content = append(f.content, data)
	return nil
}

func (f *File) Close() error {
	f.closed = true
	fmt.Printf("Closing file: %s\n", f.Name)
	return nil
}

// ProcessData accepts any Reader
func ProcessData(r Reader) {
	data, _ := r.Read()
	fmt.Printf("Processing data: %s\n", data)
}

// SaveData accepts any Writer
func SaveData(w Writer, data string) {
	w.Write(data)
	fmt.Println("Data saved")
}

func main() {
	fmt.Println("=== Interface Composition ===")

	// File implements ReadWriteCloser
	file := NewFile("config.txt")

	// Use as Writer
	SaveData(file, "host=localhost")
	SaveData(file, "port=8080")

	// Use as Reader
	ProcessData(file)

	// Use as Closer
	file.Close()

	// Demonstrate interface satisfaction
	var rw ReadWriter = file
	rw.Write("new data")
	rw.Read()

	fmt.Printf("\nFile implements multiple interfaces!\n")
}
