package main

import (
	"context"
	"fmt"
)

// Define a custom type for keys to avoid collisions with other packages.
// Using a string directly could cause accidental overwrites.
type contextKey string

const (
	requestIDKey contextKey = "requestID"
	userKey      contextKey = "user"
)

func main() {
	// Create a base context.
	ctx := context.Background()

	// Add a request ID.
	ctx = context.WithValue(ctx, requestIDKey, "abc-123")
	// Add a user object.
	ctx = context.WithValue(ctx, userKey, "context@example.com")

	// Pass the context down the call chain.
	handleRequest(ctx)
}

func handleRequest(ctx context.Context) {
	// Retrieve values using type assertion.
	if reqID, ok := ctx.Value(requestIDKey).(string); ok {
		fmt.Println("Request ID:", reqID)
	} else {
		fmt.Println("No request ID found")
	}

	if user, ok := ctx.Value(userKey).(string); ok {
		fmt.Println("User:", user)
	} else {
		fmt.Println("No user found")
	}
}
