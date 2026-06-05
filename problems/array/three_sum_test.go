package array

import (
	"reflect"
	"sort"
	"testing"
)

func normalizeTriplets(triplets [][]int) {
	for _, triplet := range triplets {
		sort.Ints(triplet)
	}

	sort.Slice(triplets, func(i, j int) bool {
		for k := 0; k < 3; k++ {
			if triplets[i][k] != triplets[j][k] {
				return triplets[i][k] < triplets[j][k]
			}
		}
		return false
	})
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			name:     "example 1",
			input:    []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "example 2",
			input:    []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			name:     "example 3",
			input:    []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			name:     "all positive",
			input:    []int{1, 2, 3, 4, 5},
			expected: [][]int{},
		},
		{
			name:     "all negative",
			input:    []int{-5, -4, -3, -2, -1},
			expected: [][]int{},
		},
		{
			name:     "multiple zeros",
			input:    []int{0, 0, 0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			name:     "single valid triplet",
			input:    []int{-2, 0, 2},
			expected: [][]int{{-2, 0, 2}},
		},
		{
			name:     "duplicate values",
			input:    []int{-2, 0, 0, 2, 2},
			expected: [][]int{{-2, 0, 2}},
		},
		{
			name:     "multiple distinct triplets",
			input:    []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4},
			expected: [][]int{
				{-4, 0, 4},
				{-4, 1, 3},
				{-4, 2, 2},
				{-2, -2, 4},
				{-2, 0, 2},
			},
		},
		{
			name:     "unsorted input",
			input:    []int{3, -2, 1, 0, -1, 2, -3},
			expected: [][]int{
				{-3, 0, 3},
				{-3, 1, 2},
				{-2, -1, 3},
				{-2, 0, 2},
				{-1, 0, 1},
			},
		},
		{
			name:     "minimum length no solution",
			input:    []int{1, 2, 3},
			expected: [][]int{},
		},
		{
			name:     "minimum length valid solution",
			input:    []int{-1, 0, 1},
			expected: [][]int{{-1, 0, 1}},
		},
		{
			name:     "large numbers",
			input:    []int{-100000, 0, 100000},
			expected: [][]int{{-100000, 0, 100000}},
		},
		{
			name:     "many duplicates",
			input:    []int{-1, -1, -1, 0, 0, 0, 1, 1, 1},
			expected: [][]int{
				{-1, 0, 1},
				{0, 0, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.input)

			normalizeTriplets(got)
			normalizeTriplets(tt.expected)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("threeSum(%v) = %v, want %v",
					tt.input, got, tt.expected)
			}
		})
	}
}
