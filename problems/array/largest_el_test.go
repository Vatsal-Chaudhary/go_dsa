package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLargest(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected int
	}{
		{
			name:     "positive numbers",
			arr:      []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "unsorted array",
			arr:      []int{10, 3, 25, 7, 1},
			expected: 25,
		},
		{
			name:     "single element",
			arr:      []int{42},
			expected: 42,
		},
		{
			name:     "negative numbers",
			arr:      []int{-10, -3, -25, -1},
			expected: -1,
		},
		{
			name:     "duplicate largest elements",
			arr:      []int{4, 9, 2, 9, 1},
			expected: 9,
		},
		{
			name:     "all same elements",
			arr:      []int{7, 7, 7, 7},
			expected: 7,
		},
		{
			name:     "empty array",
			arr:      []int{},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getLargest(tt.arr)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSecondLargest(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected int
	}{
		{
			name:     "normal case",
			arr:      []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "unsorted array",
			arr:      []int{10, 3, 25, 7, 1},
			expected: 10,
		},
		{
			name:     "duplicate largest values",
			arr:      []int{5, 5, 4, 3, 2},
			expected: 4,
		},
		{
			name:     "all same elements",
			arr:      []int{7, 7, 7, 7},
			expected: -1,
		},
		{
			name:     "negative numbers",
			arr:      []int{-10, -3, -25, -1},
			expected: -3,
		},
		{
			name:     "two elements",
			arr:      []int{8, 3},
			expected: 3,
		},
		{
			name:     "single element",
			arr:      []int{42},
			expected: -1,
		},
		{
			name:     "empty array",
			arr:      []int{},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := secondLargest(tt.arr)
			assert.Equal(t, tt.expected, result)
		})
	}
}
