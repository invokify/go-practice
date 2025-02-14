package binarysearchtree

import (
	"reflect"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Test case 1", []int{4, 2, 6, 3, 1, 5, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"Test case 2", []int{4, 2, 1, 3, 6, 5, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"Test case 3", []int{4, 6, 5, 7, 2, 1, 3}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"Test case 4", []int{4, 2, 1, 3, 6, 5, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"Test case 5", []int{4, 7, 6, 5, 2, 3, 1}, []int{1, 2, 3, 4, 5, 6, 7}},
	}

	// Test each case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := NewBst(tt.input[0])
			for _, i := range tt.input[1:] {
				bst.Insert(i)
			}

			actual := bst.SortedData()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, actual)
			}
		})
	}
}
