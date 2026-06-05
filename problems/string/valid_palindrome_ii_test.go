package string

import "testing"

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "example 1 already palindrome",
			s:    "aba",
			want: true,
		},
		{
			name: "example 2 delete one char",
			s:    "abca",
			want: true,
		},
		{
			name: "example 3 impossible",
			s:    "abc",
			want: false,
		},
		{
			name: "single character",
			s:    "a",
			want: true,
		},
		{
			name: "two equal chars",
			s:    "aa",
			want: true,
		},
		{
			name: "two different chars",
			s:    "ab",
			want: true,
		},
		{
			name: "delete left",
			s:    "deeee",
			want: true,
		},
		{
			name: "delete right",
			s:    "eeeed",
			want: true,
		},
		{
			name: "middle deletion",
			s:    "raceacar",
			want: true,
		},
		{
			name: "already palindrome even length",
			s:    "abba",
			want: true,
		},
		{
			name: "already palindrome odd length",
			s:    "racecar",
			want: true,
		},
		{
			name: "cannot fix with one deletion",
			s:    "abcdef",
			want: false,
		},
		{
			name: "tricky false",
			s:    "abcda",
			want: false,
		},
		{
			name: "long repeated",
			s:    "eeeeeeeeeeeeee",
			want: true,
		},
		{
			name: "one mismatch near center",
			s:    "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validPalindrome(tt.s)

			if got != tt.want {
				t.Errorf(
					"validPalindrome(%q) = %v, want %v",
					tt.s,
					got,
					tt.want,
				)
			}
		})
	}
}
