package main

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"hijklmmn", false},
		{"abbceffg", false},
		{"abbcegjk", false},
		{"abbcegjk", false},
		{"abcdffaa", true},
		{"ghjaabcc", true},
	}

	for _, tt := range tests {
		l := Valid(tt.in)
		if l != tt.out {
			t.Errorf("Test (%v): got %v, expected %v", tt.in, l, tt.out)
		}
	}

}

func TestFirstCandidate(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"ghijklmn", "ghjaaaaa"},
	}

	for _, tt := range tests {
		l := FirstCandidate(tt.in)
		if l != tt.out {
			t.Errorf("Test (%v): got %v, expected %v", tt.in, l, tt.out)
		}
	}

}

func TestIncrement(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"xx", "xy"},
		{"xz", "ya"},
		{"ya", "yb"},
	}

	for _, tt := range tests {
		l := Increment(tt.in)
		if l != tt.out {
			t.Errorf("Test (%v): got %v, expected %v", tt.in, l, tt.out)
		}
	}

}

func TestNext(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"abcdefgh", "abcdffaa"},
		{"ghijklmn", "ghjaabcc"},
	}

	for _, tt := range tests {
		l := NextPass(tt.in)
		if l != tt.out {
			t.Errorf("Test (%v): got %v, expected %v", tt.in, l, tt.out)
		}
	}

}
