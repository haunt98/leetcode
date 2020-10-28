package main

import (
	"testing"

	"github.com/haunt98/leetcode/pkg/compare"
)

func TestReverseString_1(t *testing.T) {
	tests := []struct {
		name string
		s    []byte
		want []byte
	}{
		{
			name: "example 1",
			s:    []byte{'h', 'e', 'l', 'l', 'o'},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			name: "example 1",
			s:    []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			want: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			reverseString_1(tc.s)
			compare.Diff(t, tc.want, tc.s)
		})
	}
}
