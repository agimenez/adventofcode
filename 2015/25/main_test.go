package main

import (
	"fmt"
	"testing"
)

func TestGetNthCode(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{1, 20151125},
		{2, 31916031},
		{3, 18749137},
		{21, 33511524},
		{16, 33071741},
	}

	for _, tt := range tests {
		l := GetNthCode(tt.in)
		if l != tt.out {
			t.Errorf("Test (%v): got %v, expected %v", tt.in, l, tt.out)
		}
	}

}

func TestGetPosFromCoords(t *testing.T) {
	tests := []struct {
		r    int
		c    int
		want int
	}{
		{1, 1, 1},
		{1, 2, 3},
		{1, 3, 6},
		{1, 4, 10},
		{1, 5, 15},
		{1, 6, 21},

		{2, 1, 2},
		{2, 2, 5},
		{2, 3, 9},
		{2, 4, 14},

		{3, 1, 4},
		{3, 2, 8},
		{3, 3, 13},

		{5, 2, 17},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("r=%d,c=%d", tt.r, tt.c)
		t.Run(name, func(t *testing.T) {
			got := GetPosFromCoords(tt.r, tt.c)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("GetPosFromCoords() = %v, want %v", got, tt.want)
			}
		})
	}
}
