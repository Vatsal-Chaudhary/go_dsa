package string

import "testing"

func TestLongestDistinctSubstring(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{
			name: "empty string",
			in:   "",
			want: 0,
		},
		{
			name: "single character",
			in:   "a",
			want: 1,
		},
		{
			name: "all unique",
			in:   "abcdef",
			want: 6,
		},
		{
			name: "all same",
			in:   "aaaaaa",
			want: 1,
		},
		{
			name: "simple repeat",
			in:   "abcabcbb",
			want: 3,
		},
		{
			name: "leetcode pwwkew",
			in:   "pwwkew",
			want: 3,
		},
		{
			name: "leetcode dvdf",
			in:   "dvdf",
			want: 3,
		},
		{
			name: "abba",
			in:   "abba",
			want: 2,
		},
		{
			name: "repeat at end",
			in:   "abcdefga",
			want: 7,
		},
		{
			name: "repeat at beginning",
			in:   "aabcdef",
			want: 6,
		},
		{
			name: "middle repeat",
			in:   "abcadef",
			want: 6,
		},
		{
			name: "alternating characters",
			in:   "abababab",
			want: 2,
		},
		{
			name: "numbers",
			in:   "123412345",
			want: 5,
		},
		{
			name: "special characters",
			in:   "!@#$%^&*!",
			want: 8,
		},
		{
			name: "spaces count as characters",
			in:   "abc def abc",
			want: 7,
		},
		{
			name: "long unique suffix",
			in:   "aaaabcdefg",
			want: 7,
		},
		{
			name: "complex case",
			in:   "tmmzuxt",
			want: 5,
		},
		{
			name: "repeating pair",
			in:   "aabbccdd",
			want: 2,
		},
		{
			name: "unique then repeat",
			in:   "abcdefghijklmnopqrstuvwxyza",
			want: 26,
		},
		{
			name: "single repeating pattern",
			in:   "xyzxyzxyz",
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestDistinctSubstring(tt.in)

			if got != tt.want {
				t.Fatalf(
					"LongestDistinctSubstring(%q) = %d, want %d",
					tt.in,
					got,
					tt.want,
				)
			}
		})
	}
}
