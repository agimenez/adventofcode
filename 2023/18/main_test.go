package main

import "testing"

func TestMain(t *testing.T) {
	p := polygon{
		{0, 0}, {2, 0}, {2, 2}, {0, 2}, {0, 0}}

	if p.Area() != 4 {
		t.Errorf("2x2 polygon failed, got %v, expected 4", p.Area())
	}

}
