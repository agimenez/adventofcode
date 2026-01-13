package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa", false},
		{"aa bb cc dd aaaa", true},
	}

	for _, tt := range tests {
		l := valid(tt.in)
		if l != tt.out {
			t.Errorf("Test (%v): got %v, expected %v", tt.in, l, tt.out)
		}
	}

}
