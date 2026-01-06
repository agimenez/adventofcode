package main

import (
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		in  string
		ops []string
		out string
	}{
		{"abcde",
			strings.Split(
				`swap position 4 with position 0
swap letter d with letter b
reverse positions 0 through 4
rotate left 1 step
move position 1 to position 4
move position 3 to position 0
rotate based on position of letter b
rotate based on position of letter d`, "\n"),
			"decab"},
	}

	for _, tt := range tests {
		l := Scramble(tt.ops, tt.in)
		if l != tt.out {
			t.Errorf("Test (%v): got %v, expected %v", tt.in, l, tt.out)
		}
	}

}

func TestApplyOp(t *testing.T) {
	tests := []struct {
		name string

		op   string
		pass string
		want string
	}{
		{"swap position", "swap position 4 with position 0", "abcde", "ebcda"},
		{"swap letter", "swap letter d with letter b", "ebcda", "edcba"},
		{"reverse", "reverse positions 0 through 4", "edcba", "abcde"},
		{"rotate immediate", "rotate left 1 step", "abcde", "bcdea"},
		{"rotate immediate", "rotate right 1 step", "abcde", "eabcd"},
		{"move", "move position 1 to position 4", "bcdea", "bdeac"},
		{"move", "move position 3 to position 0", "bdeac", "abdec"},
		{"rotate based", "rotate based on position of letter b", "abdec", "ecabd"},
		{"rotate based", "rotate based on position of letter d", "ecabd", "decab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := string(ApplyOp(tt.op, []byte(tt.pass)))
			if got != tt.want {
				t.Errorf("ApplyOp() = %v, want %v", got, tt.want)
			}
		})
	}
}
