package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"2-4,6-8", false},
		{"2-3,4-5", false},
		{"5-7,7-9", false},
		{"2-8,3-7", true},
		{"6-6,4-6", true},
		{"2-6,4-8", false},
	}

	for _, tt := range tests {
		c := fullyContains(tt.in)
		if c != tt.out {
			t.Errorf("fullyContains %q: got %v, expected %v", tt.in, c, tt.out)
		}
	}

}
