package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
		// Test cases
		tests := []struct {
			name     string
			input    []int
			key      int
			expected int
		}{
			{"Test case 1", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4},
			{"Test case 2", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9},
			{"Test case 3", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0},
			{"Test case 4", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, -1},
			{"Test case 5", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11, -1},
			{"Test case 6", []int{4, 8, 12, 16, 23, 28, 32}, 23, 4},
		}

		// Test each case
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				actual := SearchInts(tt.input, tt.key)
				if actual != tt.expected {
					t.Errorf("Expected %d, but got %d", tt.expected, actual)
				}
			})
		}
}