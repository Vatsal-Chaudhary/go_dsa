package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		start    int
		end      int
		expected int
	}{
		{
			name:     "example case",
			input:    []int{100, 180, 260, 310, 40, 535, 695},
			start:    0,
			end:      6,
			expected: 865,
		},
		{
			name:     "strictly increasing",
			input:    []int{1, 2, 3, 4, 5},
			start:    0,
			end:      4,
			expected: 4,
		},
		{
			name:     "strictly decreasing",
			input:    []int{5, 4, 3, 2, 1},
			start:    0,
			end:      4,
			expected: 0,
		},
		{
			name:     "all same values",
			input:    []int{7, 7, 7, 7},
			start:    0,
			end:      3,
			expected: 0,
		},
		{
			name:     "single element",
			input:    []int{5},
			start:    0,
			end:      0,
			expected: 0,
		},
		{
			name:     "empty array",
			input:    []int{},
			start:    0,
			end:      -1,
			expected: 0,
		},
		{
			name:     "multiple rises and falls",
			input:    []int{1, 5, 3, 8, 12},
			start:    0,
			end:      4,
			expected: 13,
		},
		{
			name:     "small fluctuations",
			input:    []int{3, 4, 2, 5, 1, 6},
			start:    0,
			end:      5,
			expected: 9,
		},
		{
			name:     "negative values",
			input:    []int{-5, -2, -10, -1},
			start:    0,
			end:      3,
			expected: 12,
		},
		{
			name:     "subarray range",
			input:    []int{10, 22, 5, 75, 65, 80},
			start:    0,
			end:      5,
			expected: 97,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProfit(tt.input, tt.start, tt.end)
			assert.Equal(t, tt.expected, result)
		})
	}
}
