package main

import (
	"testing"

	"github.com/haunt98/leetcode/pkg/compare"
)

func TestThirdMax(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{3, 2, 1},
			want: 1,
		},
		{
			name: "example 2",
			nums: []int{1, 2},
			want: 2,
		},
		{
			name: "example 3",
			nums: []int{2, 2, 3, 1},
			want: 1,
		},
		{
			name: "submission test",
			nums: []int{1, -2147483648, 2},
			want: -2147483648,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := thirdMax_1(tc.nums)
			compare.Diff(t, tc.want, got)
		})
	}
}
