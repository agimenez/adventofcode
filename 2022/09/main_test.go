package main

import (
	"testing"

	. "github.com/agimenez/adventofcode/utils"
)

func TestMove(t *testing.T) {
	tests := []struct {
		in    string
		start Point
		out   Point
	}{
		{"R 4", Point{0, 0}, Point{3, 0}},
	}

	for _, tt := range tests {
		rope := NewRope(tt.start)
		tail := rope.MoveHead(tt.in)
		if rope.Tail() != tt.out {
			t.Errorf("Move %q: got %v, expected %v", tt.in, tt.out)
		}
	}

}
