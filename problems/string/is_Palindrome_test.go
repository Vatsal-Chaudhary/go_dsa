package string

import "testing"

func TestIsPalindrome_All_Char(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "simple palindrome",
			input:    "madam",
			expected: true,
		},
		{
			name:     "simple non palindrome",
			input:    "hello",
			expected: false,
		},
		{
			name:     "single character",
			input:    "a",
			expected: true,
		},
		{
			name:     "empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "even length palindrome",
			input:    "abba",
			expected: true,
		},
		{
			name:     "odd length palindrome",
			input:    "racecar",
			expected: true,
		},
		{
			name:     "case sensitive check",
			input:    "Madam",
			expected: false,
		},
		{
			name:     "with spaces",
			input:    "nurses run",
			expected: false,
		},
		{
			name:     "numeric palindrome",
			input:    "12321",
			expected: true,
		},
		{
			name:     "numeric non palindrome",
			input:    "12345",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPalindrome_All_Chars(tt.input)

			if result != tt.expected {
				t.Errorf("isPalindrome(%q) = %v; want %v",
					tt.input,
					result,
					tt.expected,
				)
			}
		})
	}
}
