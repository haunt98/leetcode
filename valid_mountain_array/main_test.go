package main

import (
	"testing"

	"github.com/haunt98/leetcode/pkg/compare"
)

func TestValidMountainArray_1(t *testing.T) {
	tests := []struct {
		name string
		A    []int
		want bool
	}{
		{
			name: "example 1",
			A:    []int{2, 1},
			want: false,
		},
		{
			name: "example 2",
			A:    []int{3, 5, 5},
			want: false,
		},
		{
			name: "example 3",
			A:    []int{0, 3, 2, 1},
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := validMountainArray_1(tc.A)
			compare.Diff(t, tc.want, got)
		})
	}
}
