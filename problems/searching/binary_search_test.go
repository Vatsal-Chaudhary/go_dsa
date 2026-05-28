package searching

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   int
	}{
		{
			name:   "target in middle",
			arr:    []int{1, 2, 3, 4, 5},
			target: 3,
			want:   2,
		},
		{
			name:   "target at beginning",
			arr:    []int{1, 2, 3, 4, 5},
			target: 1,
			want:   0,
		},
		{
			name:   "target at end",
			arr:    []int{1, 2, 3, 4, 5},
			target: 5,
			want:   4,
		},
		{
			name:   "target not present",
			arr:    []int{1, 2, 3, 4, 5},
			target: 6,
			want:   -1,
		},
		{
			name:   "empty array",
			arr:    []int{},
			target: 10,
			want:   -1,
		},
		{
			name:   "single element found",
			arr:    []int{7},
			target: 7,
			want:   0,
		},
		{
			name:   "single element not found",
			arr:    []int{7},
			target: 3,
			want:   -1,
		},
		{
			name:   "negative numbers",
			arr:    []int{-10, -5, -2, 0, 3, 8},
			target: -2,
			want:   2,
		},
		{
			name:   "duplicates present",
			arr:    []int{1, 2, 2, 2, 3},
			target: 2,
			want:   2,
		},
		{
			name:   "target smaller than all",
			arr:    []int{5, 6, 7, 8},
			target: 1,
			want:   -1,
		},
		{
			name:   "target larger than all",
			arr:    []int{5, 6, 7, 8},
			target: 100,
			want:   -1,
		},
		{
			name:   "two elements first found",
			arr:    []int{1, 2},
			target: 1,
			want:   0,
		},
		{
			name:   "two elements second found",
			arr:    []int{1, 2},
			target: 2,
			want:   1,
		},
		{
			name:   "two elements none found",
			arr:    []int{1, 2},
			target: 3,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := binarySearch(tt.arr, tt.target)

			// For duplicate values, any valid index is acceptable
			if tt.name == "duplicates present" {
				if got < 0 || got >= len(tt.arr) || tt.arr[got] != tt.target {
					t.Errorf(
						"binarySearch(%v, %d) = %d, invalid duplicate index",
						tt.arr,
						tt.target,
						got,
					)
				}
				return
			}

			if got != tt.want {
				t.Errorf(
					"binarySearch(%v, %d) = %d, want %d",
					tt.arr,
					tt.target,
					got,
					tt.want,
				)
			}
		})
	}
}
