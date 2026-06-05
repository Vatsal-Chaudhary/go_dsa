package array

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "leetcode example 1",
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		{
			name:     "leetcode example 2",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "leetcode example 3",
			nums:     []int{5, 4, -1, 7, 8},
			expected: 23,
		},
		{
			name:     "all negative",
			nums:     []int{-8, -3, -6, -2, -5, -4},
			expected: -2,
		},
		{
			name:     "all positive",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "single negative",
			nums:     []int{-1},
			expected: -1,
		},
		{
			name:     "single positive",
			nums:     []int{10},
			expected: 10,
		},
		{
			name:     "maximum at beginning",
			nums:     []int{10, -3, -4, -5},
			expected: 10,
		},
		{
			name:     "maximum at end",
			nums:     []int{-5, -4, -3, 10},
			expected: 10,
		},
		{
			name:     "contains zero",
			nums:     []int{-2, 0, -1},
			expected: 0,
		},
		{
			name:     "alternating values",
			nums:     []int{2, -1, 2, -1, 2, -1, 2},
			expected: 5,
		},
		{
			name:     "large negative splits",
			nums:     []int{4, 5, -100, 6, 7},
			expected: 13,
		},
		{
			name:     "best subarray in middle",
			nums:     []int{-5, 4, 6, -1, 3, -10},
			expected: 12,
		},
		{
			name:     "all zeros",
			nums:     []int{0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "zero and positives",
			nums:     []int{0, 1, 2, 0, 3},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxSubArray(tt.nums)

			if got != tt.expected {
				t.Fatalf(
					"MaxSubArray(%v) = %d, want %d",
					tt.nums,
					got,
					tt.expected,
				)
			}
		})
	}
}
