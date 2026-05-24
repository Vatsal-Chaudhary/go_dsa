package array

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestMaxSumK(t *testing.T) {
    tests := []struct {
        name     string
        arr      []int
        k        int
        expected int
    }{
        {
            name:     "Example 1 - Standard case",
            arr:      []int{2, 1, 5, 1, 3, 2},
            k:        3,
            expected: 9, // 5 + 1 + 3 = 9
        },
        {
            name:     "All positive numbers",
            arr:      []int{1, 4, 2, 10, 2, 3, 1, 0, 20},
            k:        4,
            expected: 24,         
				},
        {
            name:     "With negative numbers",
            arr:      []int{1, -2, 3, 4, -5, 6, 7},
            k:        3,
            expected: 8, 
        },
        {
            name:     "k = 1 (should return max element)",
            arr:      []int{5, -3, 8, 2, -1},
            k:        1,
            expected: 8,
        },
        {
            name:     "k equals array length",
            arr:      []int{1, 2, 3, 4, 5},
            k:        5,
            expected: 15,
        },
        {
            name:     "Array with zeros",
            arr:      []int{0, 0, 5, 0, 0, 3, 0},
            k:        3,
            expected: 5,
        },
        {
            name:     "All negative numbers",
            arr:      []int{-5, -1, -3, -7, -2},
            k:        2,
            expected: -4, 
        },
        {
            name:     "Empty array",
            arr:      []int{},
            k:        3,
            expected: 0,
        },
        {
            name:     "k = 0 (edge case)",
            arr:      []int{1, 2, 3},
            k:        0,
            expected: 0,
        },
        {
            name:     "k > array length",
            arr:      []int{1, 2, 3},
            k:        5,
            expected: 0, 
        },
        {
            name:     "Large window with negatives",
            arr:      []int{10, -5, 20, -10, 30, 15},
            k:        3,
            expected: 40,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := maxSumK(tt.arr, tt.k)
            assert.Equal(t, tt.expected, result, "Failed: %s", tt.name)
        })
    }
}
