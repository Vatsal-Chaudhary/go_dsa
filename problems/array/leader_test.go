package array

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func normalize(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)

	sort.Ints(result)

	return result
}

func TestLeadersInArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "normal case",
			input:    []int{16, 17, 4, 3, 5, 2},
			expected: []int{17, 5, 2},
		},
		{
			name:     "strictly increasing array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5},
		},
		{
			name:     "strictly decreasing array",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "all same elements",
			input:    []int{7, 7, 7, 7},
			expected: []int{7},
		},
		{
			name:     "single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "negative numbers",
			input:    []int{-1, -2, -3, -4},
			expected: []int{-1, -2, -3, -4},
		},
		{
			name:     "mixed positive and negative",
			input:    []int{-1, 3, 2, -5, 1},
			expected: []int{3, 2, 1},
		},
		{
			name:     "duplicates with leaders",
			input:    []int{5, 5, 3, 2},
			expected: []int{5, 3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := leadersInArray(tt.input)

			assert.Equal(
				t,
				normalize(tt.expected),
				normalize(result),
			)
		})
	}
}
