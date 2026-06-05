package array

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "example 1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "example 2",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "all same height",
			height: []int{5, 5, 5, 5, 5},
			want:   20,
		},
		{
			name:   "strictly increasing",
			height: []int{1, 2, 3, 4, 5},
			want:   6,
		},
		{
			name:   "strictly decreasing",
			height: []int{5, 4, 3, 2, 1},
			want:   6,
		},
		{
			name:   "max in middle",
			height: []int{2, 3, 10, 5, 7, 8, 9},
			want:   36,
		},
		{
			name:   "contains zero",
			height: []int{0, 2, 0, 4, 0},
			want:   4,
		},
		{
			name:   "all zeros",
			height: []int{0, 0, 0},
			want:   0,
		},
		{
			name:   "large width wins",
			height: []int{4, 3, 2, 1, 4},
			want:   16,
		},
		{
			name:   "minimum valid input",
			height: []int{10000, 10000},
			want:   10000,
		},
		{
			name:   "two tall walls",
			height: []int{1, 100, 1, 1, 100, 1},
			want:   300,
		},
		{
			name:   "single tall wall not enough",
			height: []int{1, 1, 100, 1, 1},
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.height)

			if got != tt.want {
				t.Errorf(
					"maxArea(%v) = %d, want %d",
					tt.height,
					got,
					tt.want,
				)
			}
		})
	}
}
