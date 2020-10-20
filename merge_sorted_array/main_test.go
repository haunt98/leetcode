package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMerge_1(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{
			name:  "example",
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:  "submission test",
			nums1: []int{4, 0, 0, 0, 0, 0},
			m:     1,
			nums2: []int{1, 2, 3, 5, 6},
			n:     5,
			want:  []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			merge_1(tc.nums1, tc.m, tc.nums2, tc.n)
			if diff := cmp.Diff(tc.want, tc.nums1); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
