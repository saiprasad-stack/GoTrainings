package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a cancellable context derived from Background.
	ctx, cancel := context.WithCancel(context.Background())

	// Always call the cancel function to release resources.
	// Even if the context is cancelled automatically (e.g., by a timeout),
	// you should still call cancel() to avoid a memory leak.
	defer cancel()

	// Start a goroutine that listens for cancellation.
	go worker(ctx, "worker-1")

	// Let the worker run for a while, then cancel it.
	time.Sleep(1 * time.Second)
	fmt.Println("Main: Cancelling the context...")
	cancel()

	// Give the worker time to exit.
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Main: Done.")
}

// worker simulates a long‑running task that respects context cancellation.
func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			// ctx.Err() returns the reason for cancellation (context.Canceled or context.DeadlineExceeded)
			fmt.Printf("%s: Stopped because %v\n", name, ctx.Err())
			return
		default:
			fmt.Printf("%s: Working...\n", name)
			time.Sleep(200 * time.Millisecond)
		}
	}
}
