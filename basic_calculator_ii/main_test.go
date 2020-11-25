package main

import (
	"testing"

	"github.com/haunt98/leetcode/pkg/compare"
)

func TestParseReversePolishNotation(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{
			name: "example 1",
			s:    "3+2*2",
			want: []string{"3", "2", "2", "*", "+"},
		},
		{
			name: "example 2",
			s:    " 3/2 ",
			want: []string{"3", "2", "/"},
		},
		{
			name: "example 3",
			s:    " 3+5 / 2 ",
			want: []string{"3", "5", "2", "/", "+"},
		},
		{
			name: "testcase",
			s:    "1-1+1",
			want: []string{"1", "1", "-", "1", "+"},
		},
		{
			name: "testcase",
			s:    "1*2-3/4+5*6-7*8+9/10",
			want: []string{"1", "2", "*", "3", "4", "/", "-", "5", "6", "*", "+", "7", "8", "*", "-", "9", "10", "/", "+"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := parseReversePolishNotation(tc.s)
			compare.Diff(t, tc.want, got)
		})
	}
}

func TestAtoi(t *testing.T) {
	tests := []struct {
		name       string
		v          string
		wantResult int
		wantOK     bool
	}{
		{
			name:       "1",
			v:          "1",
			wantResult: 1,
			wantOK:     true,
		},
		{
			name:       "123",
			v:          "123",
			wantResult: 123,
			wantOK:     true,
		},
		{
			name:       "wrong",
			v:          "wrong",
			wantResult: 0,
			wantOK:     false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotResult, gotOK := atoi(tc.v)
			compare.Diff(t, tc.wantOK, gotOK)
			compare.Diff(t, tc.wantResult, gotResult)
		})
	}
}

func TestCalcReversePolishNotation(t *testing.T) {
	tests := []struct {
		name       string
		exprs      []string
		wantResult int
		wantOK     bool
	}{
		{
			name:       "example 1",
			exprs:      []string{"3", "2", "2", "*", "+"},
			wantResult: 7,
			wantOK:     true,
		},
		{
			name:       "example 2",
			exprs:      []string{"3", "2", "/"},
			wantResult: 1,
			wantOK:     true,
		},
		{
			name:       "example 3",
			exprs:      []string{"3", "5", "2", "/", "+"},
			wantResult: 5,
			wantOK:     true,
		},
		{
			name:       "testcase",
			exprs:      []string{"1", "1", "-", "1", "+"},
			wantResult: 1,
			wantOK:     true,
		},
		{
			name:       "testcase",
			exprs:      []string{"1", "2", "*", "3", "4", "/", "-", "5", "6", "*", "+", "7", "8", "*", "-", "9", "10", "/", "+"},
			wantResult: -24,
			wantOK:     true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotResult, gotOK := calcReversePolishNotation(tc.exprs)
			compare.Diff(t, tc.wantOK, gotOK)
			compare.Diff(t, tc.wantResult, gotResult)
		})
	}
}
