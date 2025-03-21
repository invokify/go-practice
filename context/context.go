package context

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type contextKey string

const (
	userIDKey contextKey = "userID"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Get the request's context
	ctx := r.Context()

	// Add a timeout to the request context
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Add a value to the context
	ctx = context.WithValue(ctx, userIDKey, r.Header.Get("User-ID"))

	// Use the context in a goroutine
	resultChan := make(chan string, 1)
	go doSlowOperation(ctx, resultChan)

	select {
	case result := <-resultChan:
		fmt.Fprintf(w, "Result: %s", result)
	case <-ctx.Done():
		http.Error(w, "Request timed out", http.StatusRequestTimeout)
	}
}

func doSlowOperation(ctx context.Context, resultChan chan<- string) {
	userID, _ := ctx.Value("userID").(string)
	select {
	case <-time.After(6 * time.Second):
		resultChan <- fmt.Sprintf("Operation completed for user %s", userID)
	case <-ctx.Done():
		fmt.Println("Operation canceled")
	}
}
