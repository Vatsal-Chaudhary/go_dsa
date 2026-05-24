package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroesToEnd(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "mixed zeroes and numbers",
			input:    []int{1, 0, 2, 0, 3, 4},
			expected: []int{1, 2, 3, 4, 0, 0},
		},
		{
			name:     "zeroes already at end",
			input:    []int{1, 2, 3, 0, 0},
			expected: []int{1, 2, 3, 0, 0},
		},
		{
			name:     "all zeroes",
			input:    []int{0, 0, 0, 0},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "no zeroes",
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "single zero",
			input:    []int{0},
			expected: []int{0},
		},
		{
			name:     "single non-zero",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "zeroes at beginning",
			input:    []int{0, 0, 1, 2, 3},
			expected: []int{1, 2, 3, 0, 0},
		},
		{
			name:     "alternating zeroes",
			input:    []int{0, 1, 0, 2, 0, 3},
			expected: []int{1, 2, 3, 0, 0, 0},
		},
		{
			name:     "negative numbers with zeroes",
			input:    []int{-1, 0, -2, 0, 3},
			expected: []int{-1, -2, 3, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := moveZeroesToEnd(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
