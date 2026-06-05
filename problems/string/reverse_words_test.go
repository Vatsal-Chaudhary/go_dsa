package string

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "two words",
			input:    "hello world",
			expected: "world hello",
		},
		{
			name:     "multiple words",
			input:    "one two three",
			expected: "three two one",
		},
		{
			name:     "single word",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "leading spaces",
			input:    " hello world",
			expected: "world hello",
		},
		{
			name:     "trailing spaces",
			input:    "hello world ",
			expected: "world hello",
		},
		{
			name:     "multiple spaces between words",
			input:    "hello   world",
			expected: "world hello",
		},
		{
			name:     "all spaces",
			input:    "     ",
			expected: "",
		},
		{
			name:     "punctuation",
			input:    "hello, world!",
			expected: "world! hello,",
		},
		{
			name:     "numbers",
			input:    "123 456 789",
			expected: "789 456 123",
		},
		{
			name:     "mixed content",
			input:    "Go 1.22 is awesome",
			expected: "awesome is 1.22 Go",
		},
		{
			name:     "unicode",
			input:    "你好 世界",
			expected: "世界 你好",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseWords(tt.input)
			if got != tt.expected {
				t.Errorf("ReverseWords(%q) = %q, want %q",
					tt.input, got, tt.expected)
			}
		})
	}
}
