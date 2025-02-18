package railfence

import "testing"

func TestDecode(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		message  string
		rails    int
		expected string
	}{
		{"Test case 1", "WECRLTEERDSOEEFEAOCAIVDEN", 3, "WEAREDISCOVEREDFLEEATONCE"},
	}

	// Test each case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Decode(tt.message, tt.rails)
			if actual != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, actual)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		message  string
		rails    int
		expected string
	}{
		{"Test case 1", "WEAREDISCOVEREDFLEEATONCE", 3, "WECRLTEERDSOEEFEAOCAIVDEN"},
	}

	// Test each case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Encode(tt.message, tt.rails)
			if actual != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, actual)
			}
		})
	}
}
