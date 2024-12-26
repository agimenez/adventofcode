package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for _, tt := range tests {
		l := floor(tt.in)
		if l != tt.out {
			t.Errorf("Test: got %v, expected %v", l, tt.out)
		}
	}

}
