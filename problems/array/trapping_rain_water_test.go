package array

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
    tests := []struct {
        name     string
        height   []int
        expected int
    }{
        {
            name:     "Example 1 - Standard case",
            height:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
            expected: 6,
        },
        {
            name:     "Example 2 - Another standard case",
            height:   []int{4, 2, 0, 3, 2, 5},
            expected: 9,
        },
        {
            name:     "No water can be trapped - increasing",
            height:   []int{1, 2, 3, 4, 5},
            expected: 0,
        },
        {
            name:     "No water can be trapped - decreasing",
            height:   []int{5, 4, 3, 2, 1},
            expected: 0,
        },
        {
            name:     "All heights equal",
            height:   []int{3, 3, 3, 3, 3},
            expected: 0,
        },
        {
            name:     "Single bar",
            height:   []int{5},
            expected: 0,
        },
        {
            name:     "Empty array",
            height:   []int{},
            expected: 0,
        },
        {
            name:     "Two bars with gap",
            height:   []int{1, 0, 1},
            expected: 1,
        },
        {
            name:     "Complex case with multiple valleys",
            height:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1, 4},
            expected: 11,
        },
        {
            name:     "Large height difference",
            height:   []int{2, 0, 2, 0, 2},
            expected: 4,
        },
        {
            name:     "Water trapped only in middle",
            height:   []int{3, 1, 4, 1, 5},
            expected: 5,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := trap(tt.height)
            assert.Equal(t, tt.expected, result, "Test case failed: %s", tt.name)
        })
    }
}
