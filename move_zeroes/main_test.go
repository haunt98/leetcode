package main

import (
	"testing"

	"github.com/haunt98/leetcode/pkg/compare"
)

func TestMoveZeroes_1(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example",
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			moveZeroes_1(tc.nums)
			compare.Diff(t, tc.want, tc.nums)
		})
	}
}
