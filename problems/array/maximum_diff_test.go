package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDifference(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "normal case",
			input:    []int{2, 3, 10, 6, 4, 8, 1},
			expected: 8, // 10 - 2
		},
		{
			name:     "strictly increasing",
			input:    []int{1, 2, 3, 4, 5},
			expected: 4, // 5 - 1
		},
		{
			name:     "strictly decreasing",
			input:    []int{9, 8, 7, 6, 5},
			expected: -1, // 8-9 or any adjacent valid pair
		},
		{
			name:     "all same elements",
			input:    []int{7, 7, 7, 7},
			expected: 0,
		},
		{
			name:     "negative numbers",
			input:    []int{-10, -3, -6, -2},
			expected: 8, // -2 - (-10)
		},
		{
			name:     "mixed positive and negative",
			input:    []int{-5, -1, 0, 3, -2, 10},
			expected: 15, // 10 - (-5)
		},
		{
			name:     "single element",
			input:    []int{42},
			expected: -1,
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: -1,
		},
		{
			name:     "max difference in middle",
			input:    []int{9, 1, 5, 3, 7},
			expected: 6, // 7 - 1
		},
		{
			name:     "duplicate values",
			input:    []int{4, 4, 4, 9},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxDifference(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
