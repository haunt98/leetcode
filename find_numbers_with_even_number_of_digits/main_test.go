package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindNumbers_1(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{12, 345, 2, 6, 7896},
			want: 2,
		},
		{
			name: "example 2",
			nums: []int{555, 901, 482, 1771},
			want: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := findNumbers_1(tc.nums)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
