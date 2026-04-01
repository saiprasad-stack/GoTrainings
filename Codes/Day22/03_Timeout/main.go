// how to set timeouts and deadlines using context.WithTimeout and context.WithDeadline.
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Example 1: WithTimeout – automatically cancels after a duration.
	// Here we set a timeout of 1 second.
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Start a goroutine that sleeps for 2 seconds, but the context will cancel after 1s.
	go func() {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Operation completed")
		case <-ctx.Done():
			fmt.Println("Timeout occurred:", ctx.Err())
		}
	}()

	// Wait a bit to see the output.
	time.Sleep(2 * time.Second)

	// Example 2: WithDeadline – cancels at an absolute time.
	deadline := time.Now().Add(2 * time.Second)
	ctx2, cancel2 := context.WithDeadline(context.Background(), deadline)
	defer cancel2()

	// Check if the deadline is set
	if d, ok := ctx2.Deadline(); ok {
		fmt.Printf("Deadline set at: %v\n", d)
	}

	// Simulate a task that takes 1 second (fits within deadline)
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Task finished before deadline")
	case <-ctx2.Done():
		fmt.Println("Task cancelled:", ctx2.Err())
	}

	// Simulate a task that takes 3 seconds (exceeds deadline)
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("This will never print because deadline is shorter")
	case <-ctx2.Done():
		fmt.Println("Deadline exceeded:", ctx2.Err())
	}
}
