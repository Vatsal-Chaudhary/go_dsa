package array

import "testing"

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{1, 7, 3, 6, 5, 6},
			want: 3,
		},
		{
			name: "example 2",
			nums: []int{1, 2, 3},
			want: -1,
		},
		{
			name: "example 3",
			nums: []int{2, 1, -1},
			want: 0,
		},
		{
			name: "single element",
			nums: []int{5},
			want: 0,
		},
		{
			name: "pivot at beginning",
			nums: []int{0, 1, -1},
			want: 0,
		},
		{
			name: "pivot at end",
			nums: []int{-1, -1, 0},
			want: -1,
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0, 0},
			want: 0,
		},
		{
			name: "multiple pivots returns leftmost",
			nums: []int{0, 0, 0},
			want: 0,
		},
		{
			name: "negative numbers",
			nums: []int{-1, -1, -1, -1, -4},
			want: -1,
		},
		{
			name: "mixed positive and negative",
			nums: []int{-1, 1, 0},
			want: 2,
		},
		{
			name: "no pivot with zeros",
			nums: []int{1, 0, 1},
			want: 1,
		},
		{
			name: "pivot in middle",
			nums: []int{2, 3, -1, 8, 4},
			want: 3,
		},
		{
			name: "alternating values",
			nums: []int{1, -1, 1, -1, 0},
			want: 4,
		},
		{
			name: "large values",
			nums: []int{1000, -1000, 1000, -1000, 0},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pivotIndex(tt.nums)
			if got != tt.want {
				t.Errorf("pivotIndex(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
