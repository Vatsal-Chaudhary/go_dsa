package array

import (
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example 1",
			nums: []int{1, 2, 3, 4},
			want: []int{1, 3, 6, 10},
		},
		{
			name: "example 2",
			nums: []int{1, 1, 1, 1, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "example 3",
			nums: []int{3, 1, 2, 10, 1},
			want: []int{3, 4, 6, 16, 17},
		},
		{
			name: "single element",
			nums: []int{5},
			want: []int{5},
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0, 0},
			want: []int{0, 0, 0, 0},
		},
		{
			name: "negative numbers",
			nums: []int{-1, -2, -3},
			want: []int{-1, -3, -6},
		},
		{
			name: "mixed positive and negative",
			nums: []int{5, -2, 3, -1},
			want: []int{5, 3, 6, 5},
		},
		{
			name: "alternating values",
			nums: []int{1, -1, 1, -1, 1},
			want: []int{1, 0, 1, 0, 1},
		},
		{
			name: "large values",
			nums: []int{1000000, 1000000, -1000000},
			want: []int{1000000, 2000000, 1000000},
		},
		{
			name: "decreasing values",
			nums: []int{5, 4, 3, 2, 1},
			want: []int{5, 9, 12, 14, 15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := runningSum(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runningSum(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
