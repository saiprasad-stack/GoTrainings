package main

import (
	"fmt"
	"strings"
)

func main() {

	// 1. Contains
	// Checks if substring exists inside string
	fmt.Println(strings.Contains("golang", "go"))   // true
	fmt.Println(strings.Contains("golang", "lang")) // true
	fmt.Println(strings.Contains("golang", "java")) // false

	// 2. HasPrefix / HasSuffix
	// Checks start and end of string
	fmt.Println(strings.HasPrefix("golang", "go")) // true
	fmt.Println(strings.HasPrefix("golang", "la")) // false

	fmt.Println(strings.HasSuffix("golang", "ng")) // true
	fmt.Println(strings.HasSuffix("golang", "go")) // false

	// 3. Index
	// Returns index of first occurrence, -1 if not found
	fmt.Println(strings.Index("golang", "l"))  // 2
	fmt.Println(strings.Index("golang", "go")) // 0
	fmt.Println(strings.Index("golang", "z"))  // -1

	// 4. Count
	// Counts number of non-overlapping occurrences
	fmt.Println(strings.Count("go go go", "go")) // 3
	fmt.Println(strings.Count("aaaa", "a"))      // 4
	fmt.Println(strings.Count("abc", "z"))       // 0

	// 5. ToUpper / ToLower
	// Converts case
	fmt.Println(strings.ToUpper("go"))     // GO
	fmt.Println(strings.ToUpper("GoLang")) // GOLANG

	fmt.Println(strings.ToLower("GO"))     // go
	fmt.Println(strings.ToLower("GoLang")) // golang

	// 6. Trim / TrimSpace
	// Trim removes specific characters, TrimSpace removes spaces
	fmt.Println(strings.Trim("!!go!!", "!"))   // go
	fmt.Println(strings.Trim("$$$go$$$", "$")) // go

	fmt.Println(strings.TrimSpace("  go  "))    // go
	fmt.Println(strings.TrimSpace("\n go \t ")) // go

	// 7. Split
	// Splits string into slice based on separator
	fmt.Println(strings.Split("a,b,c", ","))         // [a b c]
	fmt.Println(strings.Split("go-lang", "-"))       // [go lang]
	fmt.Println(strings.Split("one:two:three", ":")) // [one two three]

	// 8. Join
	// Joins slice into string with separator
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))      // a-b
	fmt.Println(strings.Join([]string{"go", "lang"}, " "))  // go lang
	fmt.Println(strings.Join([]string{"1", "2", "3"}, ",")) // 1,2,3

	// 9. Replace / ReplaceAll
	// Replace replaces limited occurrences, ReplaceAll replaces all
	fmt.Println(strings.Replace("go go", "go", "GO", 1))    // GO go
	fmt.Println(strings.Replace("go go go", "go", "GO", 2)) // GO GO go

	fmt.Println(strings.ReplaceAll("go go", "go", "GO")) // GO GO
	fmt.Println(strings.ReplaceAll("aaa", "a", "b"))     // bbb

	// 10. Repeat
	// Repeats string n times
	fmt.Println(strings.Repeat("go", 3)) // gogogo
	fmt.Println(strings.Repeat("a", 5))  // aaaaa

	// 11. Compare
	// Lexicographical comparison
	// returns:
	// 0 if equal
	// -1 if first < second
	// 1 if first > second
	fmt.Println(strings.Compare("a", "b"))   // -1
	fmt.Println(strings.Compare("b", "a"))   // 1
	fmt.Println(strings.Compare("go", "go")) // 0

	// 12. Fields
	// Splits string by whitespace
	fmt.Println(strings.Fields("go is fun"))       // [go is fun]
	fmt.Println(strings.Fields("  hello world  ")) // [hello world]
	fmt.Println(strings.Fields("a\tb\nc"))         // [a b c]
}
