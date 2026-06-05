package string

import (
	"reflect"
	"sort"
	"testing"
)

func normalizer(arr []int) []int {
	cp := append([]int(nil), arr...)
	sort.Ints(cp)
	return cp
}

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "leetcode example 1",
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "leetcode example 2",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "all elements same",
			nums:     []int{5, 5, 5, 5},
			k:        1,
			expected: []int{5},
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -1, -2, -2, -2, -3},
			k:        2,
			expected: []int{-2, -1},
		},
		{
			name:     "k equals number of unique elements",
			nums:     []int{4, 4, 1, 2, 3},
			k:        4,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "single frequent element",
			nums:     []int{7, 7, 7, 8, 9},
			k:        1,
			expected: []int{7},
		},
		{
			name:     "multiple frequencies",
			nums:     []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4},
			k:        2,
			expected: []int{3, 4},
		},
		{
			name:     "zero included",
			nums:     []int{0, 0, 0, 1, 1, 2},
			k:        2,
			expected: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TopKFrequent(tt.nums, tt.k)

			if !reflect.DeepEqual(
				normalizer(got),
				normalizer(tt.expected),
			) {
				t.Fatalf("TopKFrequent(%v, %d) = %v, want %v",
					tt.nums, tt.k, got, tt.expected)
			}
		})
	}
}
