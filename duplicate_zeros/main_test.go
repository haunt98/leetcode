package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDuplicateZeros_1(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "example 1",
			arr:  []int{1, 0, 2, 3, 0, 4, 5, 0},
			want: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			name: "example 2",
			arr:  []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "submission testcase",
			arr:  []int{0, 4, 1, 0, 0, 8, 0, 0, 3},
			want: []int{0, 0, 4, 1, 0, 0, 0, 0, 8},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			duplicateZeros_1(tc.arr)
			if diff := cmp.Diff(tc.want, tc.arr); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
