package array

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1: Mixed colors",
			input:    []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "Example 2: Small mixed",
			input:    []int{2, 0, 1},
			expected: []int{0, 1, 2},
		},
		{
			name:     "Already sorted",
			input:    []int{0, 0, 1, 1, 2, 2},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "Reverse sorted",
			input:    []int{2, 2, 1, 1, 0, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "Single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "All same colors (all 0s)",
			input:    []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "Missing one color (no 1s)",
			input:    []int{2, 0, 2, 0},
			expected: []int{0, 0, 2, 2},
		},
		{
			name:     "Empty slice",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Work on a copy to preserve the original test case structure if needed
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			sortColors(inputCopy)

			if !reflect.DeepEqual(inputCopy, tt.expected) {
				t.Errorf("sortColors(%v) got %v; want %v", tt.input, inputCopy, tt.expected)
			}
		})
	}
}
