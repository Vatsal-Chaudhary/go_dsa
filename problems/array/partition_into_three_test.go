package array

import "testing"

func TestCanPartitionIntoThree(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			name: "basic valid partition",
			arr:  []int{1, 2, 3, 0, 3},
			want: true,
		},
		{
			name: "all zeros",
			arr:  []int{0, 0, 0, 0},
			want: true,
		},
		{
			name: "negative numbers valid",
			arr:  []int{3, -1, -1, -1, 0},
			want: false,
		},
		{
			name: "sum not divisible by three",
			arr:  []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "empty array",
			arr:  []int{},
			want: false,
		},
		{
			name: "less than three elements",
			arr:  []int{1, 2},
			want: false,
		},
		{
			name: "exactly three equal elements",
			arr:  []int{2, 2, 2},
			want: true,
		},
		{
			name: "cannot partition even though divisible",
			arr:  []int{1, 1, 1, 1, 1, 1, 1},
			want: false,
		},
		{
			name: "multiple possible cuts",
			arr:  []int{0, 2, 1, -3, 3, 0, 0},
			want: false,
		},
		{
			name: "large middle partition",
			arr:  []int{1, 1, 1, 3, 1},
			want: false,
		},
		{
			name: "all negative valid",
			arr:  []int{-3, -3, -3},
			want: true,
		},
		{
			name: "single zero",
			arr:  []int{0},
			want: false,
		},
		{
			name: "three zeros",
			arr:  []int{0, 0, 0},
			want: true,
		},
		{
			name: "valid partition with repeated values",
			arr:  []int{4, 1, 0, 1, 1, 0, 1, 4},
			want: true,
		},
		{
			name: "partition at edges",
			arr:  []int{1, -1, 1, -1, 1, -1},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canPartitionIntoThree(tt.arr)

			if got != tt.want {
				t.Errorf(
					"canPartitionIntoThree(%v) = %v, want %v",
					tt.arr,
					got,
					tt.want,
				)
			}
		})
	}
}
