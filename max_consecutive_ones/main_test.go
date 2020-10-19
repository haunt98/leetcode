package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindMaxConsecutiveOnes_1(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example",
			nums: []int{1, 1, 0, 1, 1, 1},
			want: 3,
		},
		{
			name: "random",
			nums: []int{1, 0, 1, 1, 1, 1, 0, 0, 1, 0, 1},
			want: 4,
		},
		{
			name: "single 0",
			nums: []int{0},
			want: 0,
		},
		{
			name: "single 1",
			nums: []int{1},
			want: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := findMaxConsecutiveOnes_1(tc.nums)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
