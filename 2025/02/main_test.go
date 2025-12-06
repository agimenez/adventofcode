package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		in  int
		out bool
	}{
		{33333, false},
	}

	for _, tt := range tests {
		l := isValid2(tt.in)
		if l != tt.out {
			t.Errorf("Test: got %v, expected %v", l, tt.out)
		}
	}

}
