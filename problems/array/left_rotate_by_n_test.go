package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeftRotateByN(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		n        int
		expected []int
	}{
		{
			name:     "rotate by 1",
			input:    []int{1, 2, 3, 4, 5},
			n:        1,
			expected: []int{2, 3, 4, 5, 1},
		},
		{
			name:     "rotate by 2",
			input:    []int{1, 2, 3, 4, 5},
			n:        2,
			expected: []int{3, 4, 5, 1, 2},
		},
		{
			name:     "rotate by 0",
			input:    []int{1, 2, 3, 4, 5},
			n:        0,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "rotate by array length",
			input:    []int{1, 2, 3, 4, 5},
			n:        5,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "rotate by more than array length",
			input:    []int{1, 2, 3, 4, 5},
			n:        7,
			expected: []int{3, 4, 5, 1, 2},
		},
		{
			name:     "single element array",
			input:    []int{42},
			n:        3,
			expected: []int{42},
		},
		{
			name:     "empty array",
			input:    []int{},
			n:        2,
			expected: []int{},
		},
		{
			name:     "all same elements",
			input:    []int{7, 7, 7, 7},
			n:        2,
			expected: []int{7, 7, 7, 7},
		},
		{
			name:     "negative numbers",
			input:    []int{-1, -2, -3, -4},
			n:        2,
			expected: []int{-3, -4, -1, -2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := leftRotateByN(tt.input, tt.n)
			assert.Equal(t, tt.expected, result)
		})
	}
}
