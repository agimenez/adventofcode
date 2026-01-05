package main

import (
	"testing"

	"github.com/agimenez/adventofcode/utils"
)

func TestMain(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"ihgpwlah", "DDRRRD"},
		{"hijkl", ""},
		{"kglvqrro", "DDUDRLRRUDRD"},
		{"ulqzkmiv", "DRURDRUDDLLDLUURRDULRLDUUDDDRR"},
	}

	for _, tt := range tests {
		l := findShortestPath(tt.in, utils.NewPoint(0, 0), utils.NewPoint(3, 3))
		if l.Path() != tt.out {
			t.Errorf("Test (%v): got %v, expected %v", tt.in, l, tt.out)
		}
	}

}
