package main

import (
	"testing"

	"github.com/haunt98/leetcode/pkg/compare"
)

func TestCheckIfExist_1(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			name: "example 1",
			arr:  []int{10, 2, 5, 3},
			want: true,
		},
		{
			name: "example 2",
			arr:  []int{7, 1, 14, 11},
			want: true,
		},
		{
			name: "example 3",
			arr:  []int{3, 1, 7, 11},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := checkIfExist_1(tc.arr)
			compare.Diff(t, tc.want, got)
		})
	}
}
