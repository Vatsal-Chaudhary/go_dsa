package array

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "leetcode example 1",
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "leetcode example 2",
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
		{
			name:     "two elements",
			nums:     []int{3, 4},
			expected: []int{4, 3},
		},
		{
			name:     "all ones",
			nums:     []int{1, 1, 1, 1},
			expected: []int{1, 1, 1, 1},
		},
		{
			name:     "all negatives",
			nums:     []int{-1, -2, -3, -4},
			expected: []int{-24, -12, -8, -6},
		},
		{
			name:     "single zero",
			nums:     []int{1, 2, 0, 4},
			expected: []int{0, 0, 8, 0},
		},
		{
			name:     "zero at beginning",
			nums:     []int{0, 2, 3, 4},
			expected: []int{24, 0, 0, 0},
		},
		{
			name:     "zero at end",
			nums:     []int{2, 3, 4, 0},
			expected: []int{0, 0, 0, 24},
		},
		{
			name:     "two zeros",
			nums:     []int{1, 0, 3, 0},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "mixed signs",
			nums:     []int{-1, 2, -3, 4},
			expected: []int{-24, 12, -8, 6},
		},
		{
			name:     "larger values",
			nums:     []int{2, 3, 5, 7},
			expected: []int{105, 70, 42, 30},
		},
		{
			name:     "contains one",
			nums:     []int{1, 5, 10},
			expected: []int{50, 10, 5},
		},
		{
			name:     "alternating signs",
			nums:     []int{-2, 3, -4, 5},
			expected: []int{-60, 40, -30, 24},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProductExceptSelf(tt.nums)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Fatalf(
					"ProductExceptSelf(%v) = %v, want %v",
					tt.nums,
					got,
					tt.expected,
				)
			}
		})
	}
}
