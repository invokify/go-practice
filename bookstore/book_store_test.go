package bookstore

import "testing"

func TestCost(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		books    []int
		expected int
	}{
		{"one book", []int{1}, 800},
		{"twice same book", []int{1, 1}, 1600},
		{"two different books", []int{1, 2}, 800*2 - 800*2*5/100},
		{"three different books", []int{1, 2, 3}, 800*3 - 800*3*10/100},
		{"four different books", []int{1, 2, 3, 4}, 800*4 - 800*4*20/100},
		{"five different books", []int{1, 2, 3, 4, 5}, 800*5 - 800*5*25/100},
		{"groups of 4 are cheaper than 5 and 3", []int{1, 1, 2, 2, 3, 3, 4, 5}, 2 * (800*4 - 800*4*20/100)},
	}

	// Test each case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Cost(tt.books)
			if actual != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, actual)
			}
		})
	}
}
