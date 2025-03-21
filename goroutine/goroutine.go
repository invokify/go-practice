package goroutine

import (
	"context"
	"fmt"
	"time"
)

func launchWorkersWithContext() []string {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create channel for worker output
	outputChan := make(chan string, 100)
	messages := []string{}

	// Start multiple workers
	for i := 0; i < 3; i++ {
		go worker(ctx, i, outputChan)
	}

	// Simulate some work and collect messages
	done := make(chan bool)
	go func() {
		for msg := range outputChan {
			messages = append(messages, msg)
		}
		done <- true
	}()

	// Simulate some work
	time.Sleep(5 * time.Second)

	// Cancel all workers
	cancel()

	// Wait for workers to finish
	time.Sleep(1 * time.Second)
	close(outputChan)
	<-done

	return messages
}

func worker(ctx context.Context, id int, output chan<- string) {
	for {
		select {
		case <-ctx.Done():
			output <- fmt.Sprintf("Worker %d shutting down", id)
			return
		default:
			fmt.Printf("Worker %d doing work", id)
			time.Sleep(1 * time.Second)
		}
	}
}
