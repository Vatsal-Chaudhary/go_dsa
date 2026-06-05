package string

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{
			name:     "simple anagram",
			s1:       "listen",
			s2:       "silent",
			expected: true,
		},
		{
			name:     "not an anagram",
			s1:       "hello",
			s2:       "world",
			expected: false,
		},
		{
			name:     "same strings",
			s1:       "abc",
			s2:       "abc",
			expected: true,
		},
		{
			name:     "different lengths",
			s1:       "abc",
			s2:       "ab",
			expected: false,
		},
		{
			name:     "empty strings",
			s1:       "",
			s2:       "",
			expected: true,
		},
		{
			name:     "single character",
			s1:       "a",
			s2:       "a",
			expected: true,
		},
		{
			name:     "single character mismatch",
			s1:       "a",
			s2:       "b",
			expected: false,
		},
		{
			name:     "repeated characters valid",
			s1:       "aabbcc",
			s2:       "baccab",
			expected: true,
		},
		{
			name:     "repeated characters invalid",
			s1:       "aabbcc",
			s2:       "aabbc",
			expected: false,
		},
		{
			name:     "case sensitive mismatch",
			s1:       "Listen",
			s2:       "silent",
			expected: false,
		},
		{
			name:     "with spaces",
			s1:       "rail safety",
			s2:       "fairy tales",
			expected: true,
		},
		{
			name:     "numeric anagram",
			s1:       "12345",
			s2:       "54321",
			expected: true,
		},
		{
			name:     "special characters",
			s1:       "a!b@c",
			s2:       "@cb!a",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isAnagram(tt.s1, tt.s2)

			if result != tt.expected {
				t.Errorf(
					"isAnagram(%q, %q) = %v; want %v",
					tt.s1,
					tt.s2,
					result,
					tt.expected,
				)
			}
		})
	}
}
