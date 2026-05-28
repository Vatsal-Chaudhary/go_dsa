package searching

import "testing"

func TestIsPair(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   bool
	}{
		{
			name:   "pair exists in middle",
			arr:    []int{1, 2, 3, 4, 5},
			target: 7,
			want:   true,
		},
		{
			name:   "pair exists at edges",
			arr:    []int{1, 2, 3, 4, 5},
			target: 6,
			want:   true,
		},
		{
			name:   "pair does not exist",
			arr:    []int{1, 2, 3, 4, 5},
			target: 100,
			want:   false,
		},
		{
			name:   "empty array",
			arr:    []int{},
			target: 5,
			want:   false,
		},
		{
			name:   "single element",
			arr:    []int{10},
			target: 10,
			want:   false,
		},
		{
			name:   "two elements valid pair",
			arr:    []int{2, 8},
			target: 10,
			want:   true,
		},
		{
			name:   "two elements invalid pair",
			arr:    []int{2, 8},
			target: 5,
			want:   false,
		},
		{
			name:   "duplicate values valid pair",
			arr:    []int{1, 2, 2, 3, 4},
			target: 4,
			want:   true,
		},
		{
			name:   "negative numbers valid pair",
			arr:    []int{-5, -2, 0, 3, 8},
			target: 1,
			want:   true,
		},
		{
			name:   "negative numbers invalid pair",
			arr:    []int{-5, -2, 0, 3, 8},
			target: 100,
			want:   false,
		},
		{
			name:   "all zeros",
			arr:    []int{0, 0, 0, 0},
			target: 0,
			want:   true,
		},
		{
			name:   "same number cannot reuse single element",
			arr:    []int{5},
			target: 10,
			want:   false,
		},
		{
			name:   "multiple possible pairs",
			arr:    []int{1, 2, 3, 4, 5, 6},
			target: 7,
			want:   true,
		},
		{
			name:   "large values",
			arr:    []int{100, 200, 300, 400},
			target: 700,
			want:   true,
		},
		{
			name:   "target smaller than smallest possible sum",
			arr:    []int{5, 6, 7, 8},
			target: 1,
			want:   false,
		},
		{
			name:   "target larger than largest possible sum",
			arr:    []int{1, 2, 3, 4},
			target: 100,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPair(tt.arr, tt.target)

			if got != tt.want {
				t.Errorf(
					"isPair(%v, %d) = %v, want %v",
					tt.arr,
					tt.target,
					got,
					tt.want,
				)
			}
		})
	}
}
