package main

import (
	"fmt"
)

/*
========================================================
PANIC AND RECOVER IN GO
========================================================

NORMAL ERROR HANDLING IN GO
---------------------------
Most Go programs use the "error" return pattern.

Example:
    result, err := divide(a,b)
    if err != nil {
        handle error
    }

PANIC
-----
panic is used when something unexpected or impossible happens.
It immediately stops normal execution of the program.

When panic occurs:
1. Current function stops execution
2. Deferred functions execute
3. Panic moves up the call stack
4. If no recover is found, the program crashes

RECOVER
-------
recover is used to stop a panic and regain control.

IMPORTANT RULE:
recover works ONLY inside a deferred function.

If recover is not used, the program terminates.
*/

func main() {

	fmt.Println("Program Started")

	fmt.Println("\n--- Example 1 : Basic Panic ---")
	basicPanic()

	fmt.Println("\n--- Example 2 : Panic With Defer ---")
	panicWithDefer()

	fmt.Println("\n--- Example 3 : Recover Example ---")
	recoverExample()

	fmt.Println("\n--- Example 4 : Panic Propagation ---")
	panicPropagation()

	fmt.Println("\nProgram Finished")
}

/*
========================================================
EXAMPLE 1 : BASIC PANIC
========================================================

panic stops the program immediately.

Any code written after panic will NOT execute.
*/

func basicPanic() {

	fmt.Println("Before panic")

	// Uncomment this line to see program crash
	// panic("Something went wrong!")

	fmt.Println("After panic (runs only if panic is commented)")
}

/*
========================================================
EXAMPLE 2 : PANIC WITH DEFER
========================================================

When panic happens, deferred functions still run.

This is very useful for:
- closing files
- releasing locks
- cleaning resources
*/

func panicWithDefer() {

	defer fmt.Println("Deferred function executed")

	fmt.Println("Before panic")

	// Uncomment to test
	// panic("panic inside panicWithDefer")

	fmt.Println("After panic (runs only if panic disabled)")
}

/*
========================================================
EXAMPLE 3 : RECOVER
========================================================

recover stops a panic.

Rule:
recover must be used inside a deferred function.

Steps:
1. panic occurs
2. deferred function runs
3. recover catches panic
4. program continues execution
*/

func recoverExample() {

	defer func() {

		r := recover()

		if r != nil {
			fmt.Println("Recovered from panic:", r)
		}

	}()

	fmt.Println("Before panic")

	panic("Something bad happened")

	// This line will never run
	fmt.Println("After panic")
}

/*
========================================================
EXAMPLE 4 : PANIC PROPAGATION
========================================================

If panic happens inside a function and it is not recovered,
the panic travels up the call stack.

Example stack:

main -> functionA -> functionB

If panic happens in functionB:
1. functionB stops
2. functionA stops
3. main stops
4. program crashes

But if one of them recovers, execution continues.
*/

func panicPropagation() {

	defer func() {

		if r := recover(); r != nil {
			fmt.Println("Recovered in panicPropagation:", r)
		}

	}()

	functionA()

	fmt.Println("Execution continued after recovery")
}

func functionA() {
	functionB()
}

func functionB() {

	fmt.Println("Inside functionB")

	panic("Error occurred in functionB")
}

/*
========================================================
WHEN TO USE PANIC
========================================================

Good Use Cases:
1. Programmer mistakes
2. Impossible states
3. Initialization failures
4. Critical internal errors

Example:
database connection cannot be established
configuration missing

========================================================
WHEN NOT TO USE PANIC
========================================================

Do NOT use panic for normal errors like:

- file not found
- user input errors
- network errors

Instead return error values.

Example:

func readFile() (string, error)

========================================================
IMPORTANT RULES
========================================================

1. panic stops normal execution
2. defer functions still run
3. panic moves up the call stack
4. recover must be inside defer
5. if no recover exists, program crashes
*/
