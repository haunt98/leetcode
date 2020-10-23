package main

import (
	"testing"

	"github.com/haunt98/leetcode/pkg/compare"
)

func TestRemoveDuplicates_1(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		wantLen  int
		wantNums []int
	}{
		{
			name:     "example 1",
			nums:     []int{1, 1, 2},
			wantLen:  2,
			wantNums: []int{1, 2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotLen := removeDuplicates_1(tc.nums)
			compare.Diff(t, tc.wantLen, gotLen)
			compare.Diff(t, tc.wantNums, tc.nums[:gotLen])
		})
	}
}
