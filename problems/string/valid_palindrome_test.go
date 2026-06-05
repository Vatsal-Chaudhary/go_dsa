package string

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "example 1",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "example 2",
			s:    "race a car",
			want: false,
		},
		{
			name: "example 3 empty after cleanup",
			s:    " ",
			want: true,
		},
		{
			name: "simple palindrome",
			s:    "madam",
			want: true,
		},
		{
			name: "simple not palindrome",
			s:    "hello",
			want: false,
		},
		{
			name: "mixed case",
			s:    "RaceCar",
			want: true,
		},
		{
			name: "numbers palindrome",
			s:    "12321",
			want: true,
		},
		{
			name: "numbers not palindrome",
			s:    "12345",
			want: false,
		},
		{
			name: "letters and numbers",
			s:    "A1b2B1a",
			want: true,
		},
		{
			name: "symbols only",
			s:    "!!!",
			want: true,
		},
		{
			name: "skip punctuation",
			s:    ".,a,.",
			want: true,
		},
		{
			name: "apostrophe phrase",
			s:    "Madam, I'm Adam",
			want: true,
		},
		{
			name: "complex false case",
			s:    "0P",
			want: false,
		},
		{
			name: "single character",
			s:    "z",
			want: true,
		},
		{
			name: "two different chars",
			s:    "ab",
			want: false,
		},
		{
			name: "two same chars",
			s:    "aa",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.s)

			if got != tt.want {
				t.Errorf(
					"isPalindrome(%q) = %v, want %v",
					tt.s,
					got,
					tt.want,
				)
			}
		})
	}
}
