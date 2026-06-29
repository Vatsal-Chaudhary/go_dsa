package array

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		p        string
		expected []int
	}{
		{
			name:     "Example 1",
			s:        "cbaebabacd",
			p:        "abc",
			expected: []int{0, 6},
		},
		{
			name:     "Example 2",
			s:        "abab",
			p:        "ab",
			expected: []int{0, 1, 2},
		},
		{
			name:     "No anagrams",
			s:        "abcdef",
			p:        "xyz",
			expected: []int{},
		},
		{
			name:     "p longer than s",
			s:        "ab",
			p:        "abc",
			expected: []int{},
		},
		{
			name:     "All characters same",
			s:        "aaaa",
			p:        "aa",
			expected: []int{0, 1, 2},
		},
		{
			name:     "Single character",
			s:        "a",
			p:        "a",
			expected: []int{0},
		},
		{
			name:     "Empty p (edge - though constraints say >=1)",
			s:        "abc",
			p:        "",
			expected: []int{},
		},
		{
			name:     "s equals p",
			s:        "abc",
			p:        "abc",
			expected: []int{0},
		},
		{
			name:     "Overlapping anagrams",
			s:        "bacdgabcda",
			p:        "abcd",
			expected: []int{0, 5, 6},
		},
		{
			name:     "Repeated pattern",
			s:        "abababab",
			p:        "ab",
			expected: []int{0, 1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findAnagrams(tt.s, tt.p)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("findAnagrams(%q, %q) = %v, want %v", tt.s, tt.p, result, tt.expected)
			}
		})
	}
}
