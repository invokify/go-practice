package goroutine

import (
	"testing"
)

func TestLaunchWorkersWithContext(t *testing.T) {
	result := launchWorkersWithContext()

	if len(result) != 3 {
		t.Errorf("Expected 3 messages, got %d", len(result))
	}

	for _, msg := range result {
		if msg != "Worker 0 shutting down" && msg != "Worker 1 shutting down" && msg != "Worker 2 shutting down" {
			t.Errorf("Unexpected message: %s", msg)
		}
	}
}
