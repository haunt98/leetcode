package main

import (
	"testing"

	"github.com/haunt98/leetcode/pkg/compare"
)

func TestReplaceElements_1(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "example 1",
			arr:  []int{17, 18, 5, 4, 6, 1},
			want: []int{18, 6, 6, 6, 1, -1},
		},
		{
			name: "submission test",
			arr:  []int{57010, 40840, 69871, 14425, 70605},
			want: []int{70605, 70605, 70605, 70605, -1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := replaceElements_1(tc.arr)
			compare.Diff(t, tc.want, got)
		})
	}
}
