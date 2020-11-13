package main

import (
	"testing"

	"github.com/haunt98/leetcode/pkg/compare"
)

func TestKthGrammar_1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		K    int
		want int
	}{
		{
			name: "example 1",
			N:    1,
			K:    1,
			want: 0,
		},
		{
			name: "example 2",
			N:    2,
			K:    1,
			want: 0,
		},
		{
			name: "example 3",
			N:    2,
			K:    2,
			want: 1,
		},
		{
			name: "example 4",
			N:    4,
			K:    5,
			want: 1,
		},
		{
			name: "example 5",
			N:    4,
			K:    3,
			want: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := kthGrammar_1(tc.N, tc.K)
			compare.Diff(t, tc.want, got)
		})
	}
}
