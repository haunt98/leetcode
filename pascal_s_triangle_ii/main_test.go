package main

import (
	"testing"

	"github.com/haunt98/leetcode/pkg/compare"
)

func TestGetRow_1(t *testing.T) {
	tests := []struct {
		name     string
		rowIndex int
		want     []int
	}{
		{
			name:     "example 1",
			rowIndex: 3,
			want:     []int{1, 3, 3, 1},
		},
		{
			name:     "example 2",
			rowIndex: 0,
			want:     []int{1},
		},
		{
			name:     "example 3",
			rowIndex: 1,
			want:     []int{1, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := getRow_1(tc.rowIndex)
			compare.Diff(t, tc.want, got)
		})
	}
}
