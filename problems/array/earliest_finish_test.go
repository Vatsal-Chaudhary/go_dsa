package array

import "testing"

func TestEarliestFinishTime(t *testing.T) {
	tests := []struct {
		name            string
		landStartTime   []int
		landDuration    []int
		waterStartTime  []int
		waterDuration   []int
		expected        int
	}{
		{
			name:           "example 1",
			landStartTime:  []int{2, 8},
			landDuration:   []int{4, 1},
			waterStartTime: []int{6},
			waterDuration:  []int{3},
			expected:       9,
		},
		{
			name:           "example 2",
			landStartTime:  []int{5},
			landDuration:   []int{3},
			waterStartTime: []int{1},
			waterDuration:  []int{10},
			expected:       14,
		},
		{
			name:           "both open immediately",
			landStartTime:  []int{1},
			landDuration:   []int{2},
			waterStartTime: []int{1},
			waterDuration:  []int{3},
			expected:       6,
		},
		{
			name:           "must wait for water",
			landStartTime:  []int{1},
			landDuration:   []int{2},
			waterStartTime: []int{10},
			waterDuration:  []int{1},
			expected:       11,
		},
		{
			name:           "must wait for land",
			landStartTime:  []int{10},
			landDuration:   []int{1},
			waterStartTime: []int{1},
			waterDuration:  []int{2},
			expected:       11,
		},
		{
			name:           "choose second land ride",
			landStartTime:  []int{1, 5},
			landDuration:   []int{10, 1},
			waterStartTime: []int{6},
			waterDuration:  []int{1},
			expected:       7,
		},
		{
			name:           "choose second water ride",
			landStartTime:  []int{1},
			landDuration:   []int{1},
			waterStartTime: []int{1, 3},
			waterDuration:  []int{10, 1},
			expected:       4,
		},
		{
			name:           "multiple choices",
			landStartTime:  []int{2, 4},
			landDuration:   []int{3, 1},
			waterStartTime: []int{3, 8},
			waterDuration:  []int{2, 1},
			expected:       6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := earliestFinishTime(
				tt.landStartTime,
				tt.landDuration,
				tt.waterStartTime,
				tt.waterDuration,
			)

			if got != tt.expected {
				t.Fatalf("got %d, want %d", got, tt.expected)
			}
		})
	}
}
