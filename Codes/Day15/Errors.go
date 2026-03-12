package main

import (
	"errors"
	"fmt"
)

/*
====================================================
GO ERROR HANDLING
====================================================

Key Idea:
Go does NOT use exceptions like Java/Python.
Instead, functions return an error as the last value.
Caller checks error explicitly.

Why Go uses this approach:
1. Explicit error checking
2. No hidden control flow
3. Errors are just values

Other languages:
Java/Python: try { } catch { }
Go: if err != nil { }

====================================================
1. ERROR INTERFACE
====================================================

type error interface {
    Error() string
}

Any type implementing Error() string becomes an error.
*/

// simple function returning error
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

/*
====================================================
2. SENTINEL ERRORS
====================================================

A sentinel error is a predefined error value that acts as a constant.
Callers compare the returned error against this value.

Example: io.EOF, sql.ErrNoRows.

Advantages:
- Simple, direct check for a specific condition.

Disadvantages:
- Creates tight coupling between packages.
- Becomes part of your public API.
- Prevents adding context without breaking the equality check.
- Should be used rarely and only for very stable, low-level cases.

Modern Go inspection: use errors.Is to unwrap and compare.
*/

// ErrNotFound is a sentinel error indicating that an item does not exist.
var ErrNotFound = errors.New("item not found")

func findItem(id int) error {
	if id != 42 {
		return ErrNotFound
	}
	return nil
}

/*
====================================================
3. CUSTOM ERROR TYPES
====================================================

A custom error type is a struct that implements the error interface.
It can carry additional structured information.

Advantages:
- Can provide more context than a simple value.
- Callers can use type assertions to access the extra fields.

Disadvantages:
- Still creates strong coupling because the caller must import your type.
- Makes APIs brittle.

Modern Go inspection: use errors.As to extract the custom type.
*/

// MyError is a custom error type with a code and message.
type MyError struct {
	Code int
	Msg  string
}

func (e MyError) Error() string {
	return fmt.Sprintf("code %d: %s", e.Code, e.Msg)
}

func doSomethingRisky(flag bool) error {
	if !flag {
		return MyError{Code: 400, Msg: "invalid flag"}
	}
	return nil
}

/*
====================================================
4. OPAQUE ERRORS
====================================================

An opaque error is returned without any assumptions about its content.
The caller only knows that an operation succeeded or failed.

Advantages:
- Maximum flexibility for the function author.
- Minimal coupling between packages.
- This should be your default error handling strategy.

Example: most functions in the standard library return opaque errors.
*/

func processData(input string) error {
	if input == "" {
		// Return an opaque error: we don't expose what went wrong.
		return errors.New("process failed")
	}
	return nil
}

/*
====================================================
5. WRAPPING ERRORS
====================================================

Wrapping adds context to an error while preserving the original error.
This allows you to annotate the error as it bubbles up.

Go provides fmt.Errorf with the %w verb for wrapping.
The original error can be retrieved later using errors.Unwrap,
errors.Is, or errors.As.
*/

func loadConfig() error {
	err := findItem(99) // returns ErrNotFound
	if err != nil {
		// Wrap the error with context.
		return fmt.Errorf("load config: %w", err)
	}
	return nil
}

/*
====================================================
6. ASSERTING FOR BEHAVIOR, NOT TYPE
====================================================

When you need to inspect an error, prefer to check its behavior
(what it can do) rather than its concrete type or value.
Define a small interface that captures the behavior you care about.

*/

/*
====================================================
7. HANDLE ERRORS ONCE (NEVER LOG AND RETURN)
====================================================

If you log an error and also return it to the caller, you create
duplicate log entries and lose context. Instead, either handle it
(log and recover) or return it (possibly wrapped). Never both.
*/

// BadPractice logs AND returns the error – DO NOT DO THIS.
func BadPractice() error {
	err := doSomethingRisky(false)
	if err != nil {
		fmt.Println("logged:", err) // BAD: logging here
		return err                  // BAD: also returning
	}
	return nil
}

// GoodPractice wraps and returns the error – correct.
func GoodPractice() error {
	err := doSomethingRisky(false)
	if err != nil {
		// Wrap with context and return. Logging happens at the top level.
		return fmt.Errorf("GoodPractice failed: %w", err)
	}
	return nil
}

/*
====================================================
8. PANIC VS ERROR
====================================================

ERROR:
- Expected problems (file not found, invalid input).
- Program continues running.

PANIC:
- Unrecoverable problems (nil pointer dereference, programmer bug).
- Program crashes unless recovered.
*/

func panicExample() {
	panic("something went terribly wrong")
}

/*
====================================================
9. RECOVER PATTERN
====================================================

recover() catches a panic and prevents the program from crashing.
It only works inside a deferred function.
*/

func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	panic("panic inside safeFunction")
}

/*
====================================================
10. COMMA OK IDIOM
====================================================

The comma ok idiom is used in several places to indicate success:

1. Map lookup:   value, ok := myMap[key]
2. Type assertion: value, ok := i.(Type)
3. Channel receive: value, ok := <-ch

It makes success explicit and avoids panics.
*/

func commaOk() {
	// Map lookup
	scores := map[string]int{"Alice": 90}
	score, ok := scores["Bob"]
	if !ok {
		fmt.Println("Bob not found in map")
	} else {
		fmt.Println("Bob's score:", score)
	}

	// Type assertion
	var i interface{} = "hello"
	str, ok := i.(string)
	if ok {
		fmt.Println("String value:", str)
	}

}

/*
====================================================
MAIN FUNCTION
====================================================
*/

func main() {
	fmt.Println("=== 1.")
	res, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}

	fmt.Println("\n=== 2.")
	err = findItem(99)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Item not found (as expected)")
	} else {
		fmt.Println("Unexpected error:", err)
	}

	fmt.Println("\n=== 3.")
	err = doSomethingRisky(false)
	var myErr MyError
	if errors.As(err, &myErr) {
		fmt.Printf("Custom error caught: code=%d msg=%s\n", myErr.Code, myErr.Msg)
	}

	fmt.Println("\n=== 4.")
	err = processData("")
	if err != nil {
		fmt.Println("Opaque error:", err)
	}

	fmt.Println("\n=== 5.")
	err = loadConfig()
	if err != nil {
		fmt.Println("Wrapped error:", err)
		// We can still detect the underlying sentinel.
		if errors.Is(err, ErrNotFound) {
			fmt.Println("  Not Found")
		}
	}

	fmt.Println("\n=== 7.")
	fmt.Println("Calling GoodPractice (wraps and returns):")
	err = GoodPractice()
	if err != nil {
		fmt.Println("Top-level log:", err)
	}

	// fmt.Println("\nCalling BadPractice (logs and returns):")
	// err = BadPractice()
	// if err != nil {
	// 	fmt.Println("Top-level log again:", err) // duplicate
	// }

	fmt.Println("\n=== 8 & 9.")
	safeFunction()
	fmt.Println("Program continues after recover")

	fmt.Println("\n=== 10.")
	commaOk()
}
