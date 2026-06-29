package array

import (
	"testing"
)

func TestNumArray_SumRange(t *testing.T) {
	// Initialize the NumArray with the example array from LeetCode
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)

	// Define the table of test cases
	tests := []struct {
		name  string
		left  int
		right int
		want  int
	}{
		{
			name:  "Example Case 1: Elements 0 to 2",
			left:  0,
			right: 2,
			want:  1, // (-2) + 0 + 3
		},
		{
			name:  "Example Case 2: Elements 2 to 5",
			left:  2,
			right: 5,
			want:  -1, // 3 + (-5) + 2 + (-1)
		},
		{
			name:  "Example Case 3: Elements 0 to 5",
			left:  0,
			right: 5,
			want:  -3, // (-2) + 0 + 3 + (-5) + 2 + (-1)
		},
		{
			name:  "Edge Case: Single element at index 0",
			left:  0,
			right: 0,
			want:  -2,
		},
		{
			name:  "Edge Case: Single element at index 3",
			left:  3,
			right: 3,
			want:  -5,
		},
	}

	// Loop through and run the test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := obj.SumRange(tt.left, tt.right)
			if got != tt.want {
				t.Errorf("NumArray.SumRange(%d, %d) = %d; want %d", tt.left, tt.right, got, tt.want)
			}
		})
	}
}
