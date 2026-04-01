package main

/*
UNIT TESTING THEORY

Unit testing means testing a single unit of code (usually a function)
in isolation.

Goal:
- Verify correctness of smallest piece of logic
- Ensure function behaves as expected for given input

Logical reasoning:
If input is known and output is predictable,
we can validate correctness using comparison.

Example:
Add(2,3) should always return 5
*/

func Add(a, b int) int {
	return a + b
}
