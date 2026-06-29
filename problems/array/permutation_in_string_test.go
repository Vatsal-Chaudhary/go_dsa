package array

import "testing"

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{
			name: "Example 1",
			s1:   "ab",
			s2:   "eidbaooo",
			want: true,
		},
		{
			name: "Example 2",
			s1:   "ab",
			s2:   "eidboaoo",
			want: false,
		},
		{
			name: "Single Character Match",
			s1:   "a",
			s2:   "a",
			want: true,
		},
		{
			name: "Single Character No Match",
			s1:   "a",
			s2:   "b",
			want: false,
		},
		{
			name: "Exact Match",
			s1:   "abc",
			s2:   "abc",
			want: true,
		},
		{
			name: "Permutation At Beginning",
			s1:   "abc",
			s2:   "cbaxyz",
			want: true,
		},
		{
			name: "Permutation In Middle",
			s1:   "abc",
			s2:   "xxcabyy",
			want: true,
		},
		{
			name: "Permutation At End",
			s1:   "abc",
			s2:   "xyzbca",
			want: true,
		},
		{
			name: "Repeated Characters Match",
			s1:   "aabc",
			s2:   "eidcaabooo",
			want: true,
		},
		{
			name: "Repeated Characters No Match",
			s1:   "aabc",
			s2:   "abcabc",
			want: true,
		},
		{
			name: "s1 Longer Than s2",
			s1:   "abcd",
			s2:   "abc",
			want: false,
		},
		{
			name: "No Common Characters",
			s1:   "abc",
			s2:   "dddddddd",
			want: false,
		},
		{
			name: "Multiple Possible Permutations",
			s1:   "adc",
			s2:   "dcda",
			want: true,
		},
		{
			name: "Many Repeated Characters",
			s1:   "aaa",
			s2:   "aaaaaa",
			want: true,
		},
		{
			name: "Repeated Characters Insufficient",
			s1:   "aaaa",
			s2:   "aaabaaa",
			want: false,
		},
		{
			name: "Window Slides Many Times",
			s1:   "xyz",
			s2:   "aaaaaaaaaxyzb",
			want: true,
		},
		{
			name: "Permutation Not Present",
			s1:   "xyz",
			s2:   "abcdefghijkl",
			want: false,
		},
		{
			name: "Duplicate Characters Exact",
			s1:   "abb",
			s2:   "eidbaboo",
			want: true,
		},
		{
			name: "Duplicate Characters Wrong Counts",
			s1:   "abb",
			s2:   "eidabaoo",
			want: false,
		},
		{
			name: "Entire String Is Permutation",
			s1:   "listen",
			s2:   "silent",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkInclusion(tt.s1, tt.s2)
			if got != tt.want {
				t.Errorf("checkInclusion(%q, %q) = %v, want %v",
					tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}
