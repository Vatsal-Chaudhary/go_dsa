package string

import (
	"reflect"
	"sort"
	"testing"
)

func normalize(v []int) []int {
	if v == nil {
		return []int{}
	}

	res := make([]int, len(v))
	copy(res, v)
	sort.Ints(res)
	return res
}

func TestSearchPattern(t *testing.T) {
	tests := []struct {
		name    string
		txt     string
		pattern string
		want    []int
	}{
		{
			name:    "single occurrence",
			txt:     "hello world",
			pattern: "world",
			want:    []int{6},
		},
		{
			name:    "multiple occurrences",
			txt:     "abcabcabc",
			pattern: "abc",
			want:    []int{0, 3, 6},
		},
		{
			name:    "overlapping occurrences",
			txt:     "aaaaa",
			pattern: "aa",
			want:    []int{0, 1, 2, 3},
		},
		{
			name:    "pattern not found",
			txt:     "hello world",
			pattern: "xyz",
			want:    []int{},
		},
		{
			name:    "pattern equals text",
			txt:     "golang",
			pattern: "golang",
			want:    []int{0},
		},
		{
			name:    "pattern at beginning",
			txt:     "test case example",
			pattern: "test",
			want:    []int{0},
		},
		{
			name:    "pattern at end",
			txt:     "test case example",
			pattern: "example",
			want:    []int{10},
		},
		{
			name:    "single character pattern",
			txt:     "banana",
			pattern: "a",
			want:    []int{1, 3, 5},
		},
		{
			name:    "repeated character text",
			txt:     "bbbbbb",
			pattern: "bb",
			want:    []int{0, 1, 2, 3, 4},
		},
		{
			name:    "case sensitive",
			txt:     "Go go GO",
			pattern: "go",
			want:    []int{3},
		},
		{
			name:    "with spaces",
			txt:     "the cat sat on the cat",
			pattern: "cat",
			want:    []int{4, 19},
		},
		{
			name:    "unicode text",
			txt:     "你好世界你好",
			pattern: "你好",
			want:    []int{0, 4},
		},
		{
			name:    "empty text",
			txt:     "",
			pattern: "a",
			want:    []int{},
		},
		{
			name:    "empty pattern",
			txt:     "abc",
			pattern: "",
			want:    []int{},
		},
		{
			name:    "both empty",
			txt:     "",
			pattern: "",
			want:    []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchPattern(tt.txt, tt.pattern)

			if !reflect.DeepEqual(normalize(got), normalize(tt.want)) {
				t.Fatalf(
					"SearchPattern(%q, %q) = %v, want %v",
					tt.txt,
					tt.pattern,
					got,
					tt.want,
				)
			}
		})
	}
}
