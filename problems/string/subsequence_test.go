package string

import "testing"

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{
			name:     "valid subsequence",
			s1:       "abc",
			s2:       "aebdc",
			expected: true,
		},
		{
			name:     "not a subsequence",
			s1:       "aec",
			s2:       "abcde",
			expected: false,
		},
		{
			name:     "empty subsequence",
			s1:       "",
			s2:       "abc",
			expected: true,
		},
		{
			name:     "both empty",
			s1:       "",
			s2:       "",
			expected: true,
		},
		{
			name:     "larger string cannot be subsequence",
			s1:       "abcdef",
			s2:       "abc",
			expected: false,
		},
		{
			name:     "same strings",
			s1:       "golang",
			s2:       "golang",
			expected: true,
		},
		{
			name:     "repeated characters valid",
			s1:       "aab",
			s2:       "aaabbb",
			expected: true,
		},
		{
			name:     "repeated characters invalid",
			s1:       "aaa",
			s2:       "aab",
			expected: false,
		},
		{
			name:     "case sensitive mismatch",
			s1:       "ABC",
			s2:       "abc",
			expected: false,
		},
		{
			name:     "characters in wrong order",
			s1:       "cba",
			s2:       "abc",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSubsequence(tt.s1, tt.s2)

			if result != tt.expected {
				t.Errorf(
					"isSubsequence(%q, %q) = %v; want %v",
					tt.s1,
					tt.s2,
					result,
					tt.expected,
				)
			}
		})
	}
}
