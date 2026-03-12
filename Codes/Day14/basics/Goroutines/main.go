package main

import "fmt"

func main() {
	defer fmt.Println("This is third defer")
	defer fmt.Println("This is deferred")
	for i := 0; i <= 6; i++ {
		fmt.Println("Number :", i)
	}
	defer fmt.Println("This is another defer")
	fmt.Println("Goroutines")

	// Add(0)
	// switch
	// sync -- WaitGroups
	// Wait()
	// Done()

	// Defer --
	// defer file.close()

	// Multiple defer

}
