package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		in  []string
		out int
	}{
		{
			[]string{
				"...0...",
				"...1...",
				"...2...",
				"6543456",
				"7.....7",
				"8.....8",
				"9.....9",
			}, 2,
		},
		{
			[]string{
				"89010123",
				"78121874",
				"87430965",
				"96549874",
				"45678903",
				"32019012",
				"01329801",
				"10456732",
			}, 36,
		},
	}

	for _, tt := range tests {
		l := countScores(tt.in)
		if l != tt.out {
			t.Errorf("Test: got %v, expected %v", l, tt.out)
		}
	}

}
