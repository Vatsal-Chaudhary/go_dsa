package array

import "testing"

func TestEquilibrium(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			name: "basic equilibrium exists",
			arr:  []int{1, 3, 5, 2, 2},
			want: true,
		},
		{
			name: "no equilibrium",
			arr:  []int{1, 2, 3},
			want: false,
		},
		{
			name: "single element",
			arr:  []int{10},
			want: true,
		},
		{
			name: "two elements no equilibrium",
			arr:  []int{1, 2},
			want: false,
		},
		{
			name: "equilibrium at first index",
			arr:  []int{0, -3, 3},
			want: true,
		},
		{
			name: "equilibrium at last index",
			arr:  []int{-3, 3, 0},
			want: true,
		},
		{
			name: "negative numbers",
			arr:  []int{-7, 1, 5, 2, -4, 3, 0},
			want: true,
		},
		{
			name: "all zeros",
			arr:  []int{0, 0, 0, 0},
			want: true,
		},
		{
			name: "empty array",
			arr:  []int{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := equilibrium(tt.arr)

			if got != tt.want {
				t.Errorf("equilibrium(%v) = %v, want %v",
					tt.arr, got, tt.want)
			}
		})
	}
}
