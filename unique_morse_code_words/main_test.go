package main

import (
	"testing"

	"github.com/haunt98/leetcode/pkg/compare"
)

func TestUniqueMorseRepresentations_1(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  int
	}{
		{
			name:  "example",
			words: []string{"gin", "zen", "gig", "msg"},
			want:  2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := uniqueMorseRepresentations_1(tc.words)
			compare.Diff(t, tc.want, got)
		})
	}
}
