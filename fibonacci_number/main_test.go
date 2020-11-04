package main

import (
	"testing"

	"github.com/haunt98/leetcode/pkg/compare"
)

func TestFib_1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{
			name: "example 1",
			N:    1,
			want: 1,
		},
		{
			name: "example 2",
			N:    3,
			want: 2,
		},
		{
			name: "example 3",
			N:    4,
			want: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := fib_1(tc.N)
			compare.Diff(t, tc.want, got)
		})
	}
}

func TestFib_2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{
			name: "example 1",
			N:    1,
			want: 1,
		},
		{
			name: "example 2",
			N:    3,
			want: 2,
		},
		{
			name: "example 3",
			N:    4,
			want: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := fib_2(tc.N)
			compare.Diff(t, tc.want, got)
		})
	}
}
