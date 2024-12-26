package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		in     string
		out    int
		ribbon int
	}{
		{"2x3x4", 58, 34},
		{"1x1x10", 43, 14},
	}

	for _, tt := range tests {
		l, r := sqFeet(tt.in)
		if l != tt.out {
			t.Errorf("Test: got %v, expected %v", l, tt.out)
		}
		if r != tt.ribbon {
			t.Errorf("ribbon: got %v, expected %v", r, tt.ribbon)
		}
	}

}
