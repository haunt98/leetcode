package main

import (
	"testing"

	"github.com/haunt98/leetcode/pkg/compare"
)

func TestSmallerNumbersThanCurrent_1(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example 1",
			nums: []int{8, 1, 2, 2, 3},
			want: []int{4, 0, 1, 1, 3},
		},
		{
			name: "example 2",
			nums: []int{6, 5, 4, 8},
			want: []int{2, 1, 0, 3},
		},
		{
			name: "",
			nums: []int{7, 7, 7, 7},
			want: []int{0, 0, 0, 0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := smallerNumbersThanCurrent_1(tc.nums)
			compare.Diff(t, tc.want, got)
		})
	}
}

func TestSmallerNumbersThanCurrent_2(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example 1",
			nums: []int{8, 1, 2, 2, 3},
			want: []int{4, 0, 1, 1, 3},
		},
		{
			name: "example 2",
			nums: []int{6, 5, 4, 8},
			want: []int{2, 1, 0, 3},
		},
		{
			name: "",
			nums: []int{7, 7, 7, 7},
			want: []int{0, 0, 0, 0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := smallerNumbersThanCurrent_2(tc.nums)
			compare.Diff(t, tc.want, got)
		})
	}
}
