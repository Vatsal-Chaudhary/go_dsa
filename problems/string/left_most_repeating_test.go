package string

import "testing"

func TestLeftMostRepeatingChar(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "simple repeating",
			input:    "geeksforgeeks",
			expected: 0,
		},
		{
			name:     "no repeating characters",
			input:    "abcd",
			expected: -1,
		},
		{
			name:     "middle character repeats first",
			input:    "abccbd",
			expected: 1,
		},
		{
			name:     "all same characters",
			input:    "aaaa",
			expected: 0,
		},
		{
			name:     "single character",
			input:    "a",
			expected: -1,
		},
		{
			name:     "empty string",
			input:    "",
			expected: -1,
		},
		{
			name:     "last character repeats earlier",
			input:    "abcda",
			expected: 0,
		},
		{
			name:     "case sensitive",
			input:    "aAbBcCaa",
			expected: 0,
		},
		{
			name:     "space repeats",
			input:    "ab cd e ",
			expected: 2,
		},
		{
			name:     "numeric characters",
			input:    "123451",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := leftMostRepeatingChar(tt.input)

			if result != tt.expected {
				t.Errorf(
					"leftMostRepeatingChar(%q) = %d; want %d",
					tt.input,
					result,
					tt.expected,
				)
			}
		})
	}
}
