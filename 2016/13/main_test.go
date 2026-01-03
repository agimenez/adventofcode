package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", 16},
	}

	for _, tt := range tests {
		l := len(tt.in)
		if l != tt.out {
			t.Errorf("Test (%v): got %v, expected %v", tt.in, l, tt.out)
		}
	}

}
