package main

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		input     string
		expected  int
		expected2 int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 71, 231},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51, 51},
		{"2 * 3 + (4 * 5)", 26, 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437, 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240, 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632, 23340},
		{"(2 * 5) * 2 * (4 + 7 + 2) + 7", 13632, 400},
		{"(7 + 6 + (8 + 3) + 5 + 8 * 2) * 3", 13632, 222},
	}

	for _, tt := range tests {
		v := SolveExpression(tt.input, Solve)
		if v != tt.expected {
			t.Errorf("Got %d, expected %d Solve '%s'", v, tt.expected, tt.input)
		}

		v = SolveExpression(tt.input, Solve2)
		if v != tt.expected2 {
			t.Errorf("Got %d, expected %d Solve2 '%s'", v, tt.expected2, tt.input)
		}
	}
}
