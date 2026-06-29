package array

import "testing"

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "Example 1",
			s:    "ABAB",
			k:    2,
			want: 4,
		},
		{
			name: "Example 2",
			s:    "AABABBA",
			k:    1,
			want: 4,
		},
		{
			name: "Single Character",
			s:    "A",
			k:    0,
			want: 1,
		},
		{
			name: "Already Same",
			s:    "AAAA",
			k:    2,
			want: 4,
		},
		{
			name: "No Replacements Allowed",
			s:    "ABCDE",
			k:    0,
			want: 1,
		},
		{
			name: "Replace Everything",
			s:    "ABCDE",
			k:    4,
			want: 5,
		},
		{
			name: "Alternating Characters",
			s:    "ABABAB",
			k:    2,
			want: 5,
		},
		{
			name: "Large Block",
			s:    "BAAAB",
			k:    2,
			want: 5,
		},
		{
			name: "Shrink Window Required",
			s:    "ABAA",
			k:    0,
			want: 2,
		},
		{
			name: "Complex Case",
			s:    "ABBB",
			k:    2,
			want: 4,
		},
		{
			name: "Another Complex Case",
			s:    "BAAABABB",
			k:    2,
			want: 6,
		},
		{
			name: "Long Same Prefix",
			s:    "AAAAABBC",
			k:    1,
			want: 6,
		},
		{
			name: "Window Must Shrink Multiple Times",
			s:    "ABCABC",
			k:    1,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := characterReplacement(tt.s, tt.k)
			if got != tt.want {
				t.Errorf("characterReplacement(%q, %d) = %d, want %d",
					tt.s, tt.k, got, tt.want)
			}
		})
	}
}
