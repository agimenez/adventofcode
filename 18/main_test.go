package main

import "testing"

func TestGame(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 71},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}

	for _, tt := range tests {
		v := SolveExpression(tt.input)
		if v != tt.expected {
			t.Errorf("Got %d, expected %d", v, tt.expected)
		}
	}
}
