package main

import (
	"fmt"
	"path/filepath"
)

/*
filepath PACKAGE (Path handling)

Methods:
- Join
  Joins multiple path elements safely

- Base
  Returns the last element of path

- Dir
  Returns directory part of path

- Ext
  Returns file extension

- Clean
  Cleans and normalizes path
*/

func main() {
	path := filepath.Join("dir", "file.txt")
	fmt.Println(path)
}
