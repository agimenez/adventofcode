package main

import (
	"testing"

	"github.com/agimenez/adventofcode/utils"
)

func TestMain(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 2},
		{6, 1},
		{7, 2},
		{8, 1},
		{9, 2},

		{10, 3},
		{11, 2},
		{12, 3},
		{13, 4},
		{17, 4},
		{22, 3},
		{23, 2},
		{45, 4},
		{50, 7},
		{51, 6},

		{1024, 31},
	}

	sp := genSpiral(1024)
	for _, tt := range tests {
		p := sp[tt.in]

		dist := p.ManhattanDistance(utils.P0)

		if dist != tt.out {
			t.Errorf("Test (%v): got %v, expected %v", tt.in, dist, tt.out)
		}
	}

}
