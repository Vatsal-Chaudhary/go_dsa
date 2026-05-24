package array

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSubarraySumExists(t *testing.T) {
    tests := []struct {
        name     string
        arr      []int
        target   int
        expected bool
    }{
        {
            name:     "Example 1 - Standard case",
            arr:      []int{1, 4, 20, 3, 10, 5},
            target:   33,
            expected: true, // 20 + 3 + 10 = 33
        },
        {
            name:     "Subarray at beginning",
            arr:      []int{5, 3, 2, 7, 1},
            target:   8,
            expected: true, // 5 + 3 = 8
        },
        {
            name:     "Subarray at end",
            arr:      []int{1, 2, 3, 4, 5},
            target:   9,
            expected: true, // 4 + 5 = 9
        },
        {
            name:     "With zeros",
            arr:      []int{1, 0, 0, 3, 0},
            target:   0,
            expected: true,
        },
        {
            name:     "No such subarray",
            arr:      []int{1, 2, 3, 4},
            target:   10,
            expected: false,
        },
        {
            name:     "Empty array",
            arr:      []int{},
            target:   0,
            expected: false,
        },
        {
            name:     "Array with one element",
            arr:      []int{5},
            target:   5,
            expected: true,
        },
        {
            name:     "Array with one element - not match",
            arr:      []int{5},
            target:   10,
            expected: false,
        },
        {
            name:     "Target zero with all positive numbers",
            arr:      []int{1, 2, 3},
            target:   0,
            expected: false,
        },
        {
            name:     "Large numbers",
            arr:      []int{100, -50, 60, -20, 30},
            target:   70,
            expected: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := subarraySumExists(tt.arr, tt.target)
            assert.Equal(t, tt.expected, result, "Failed: %s", tt.name)
        })
    }
}
