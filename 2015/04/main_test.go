package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	}

	for _, tt := range tests {
		l := hackHash(tt.in)
		if l != tt.out {
			t.Errorf("Test: got %v, expected %v", l, tt.out)
		}
	}

}
