package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"2x3x4", 58},
		{"1x1x10", 43},
	}

	for _, tt := range tests {
		l := sqFeet(tt.in)
		if l != tt.out {
			t.Errorf("Test: got %v, expected %v", l, tt.out)
		}
	}

}
