package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"test", 4},
	}

	for _, tt := range tests {
		l := len(tt.in)
		if l != tt.out {
			t.Errorf("Test (%v): got %v, expected %v", tt.in, l, tt.out)
		}
	}

}

func TestMove(t *testing.T) {
	tests := []struct {
		progs string
		move  string
		want  string
	}{
		{"abcde", "s1", "eabcd"},
		{"eabcd", "x3/4", "eabdc"},
		{"eabdc", "pe/b", "baedc"},
	}
	for _, tt := range tests {
		t.Run(tt.move, func(t *testing.T) {
			got := string(Move([]byte(tt.progs), tt.move))
			if got != tt.want {
				t.Errorf("Move() = %v, want %v", got, tt.want)
			}
		})
	}
}
