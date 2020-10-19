package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSortedSquares_1(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "example 1",
			arr:  []int{-4, -1, 0, 3, 10},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "example 2",
			arr:  []int{-7, -3, 2, 3, 11},
			want: []int{4, 9, 9, 49, 121},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := sortedSquares_1(tc.arr)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "simple",
			arr:  []int{4, 3, 2, 10, 12, 1, 5, 6},
			want: []int{1, 2, 3, 4, 5, 6, 10, 12},
		},
		{
			name: "increasing",
			arr:  []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			insertionSort(tc.arr)
			if diff := cmp.Diff(tc.want, tc.arr); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
