package main

import (
	"testing"

	"github.com/agimenez/adventofcode/utils"
)

func TestMain(t *testing.T) {
	tests := []struct {
		in  string
		out string

		max int
	}{
		{"ihgpwlah", "DDRRRD", 370},
		{"hijkl", "", 0},
		{"kglvqrro", "DDUDRLRRUDRD", 492},
		{"ulqzkmiv", "DRURDRUDDLLDLUURRDULRLDUUDDDRR", 830},
	}

	for _, tt := range tests {
		short, long := findShortestPath(tt.in, utils.NewPoint(0, 0), utils.NewPoint(3, 3))
		if short.Path() != tt.out {
			t.Errorf("Test (%v): got %v, expected %v", tt.in, short.Path(), tt.out)
		}

		if long.cost != tt.max {
			t.Errorf("Test (%v): got %v, expected %v", tt.in, long.cost, tt.max)
		}
	}

}
