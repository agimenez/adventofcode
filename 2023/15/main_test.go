package main

import "testing"

func TestHash(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"HASH", 52},
		{"rn=1", 30},
		{"cm-", 253},
		{"qp=3", 97},
		{"qp-", 14},
		{"pc=4", 180},
		{"ot=9", 9},
		{"ab=5", 197},
		{"pc-", 48},
		{"pc=6", 214},
		{"ot=7", 231},
	}

	for _, tt := range tests {
		h := hash(tt.in)
		if h != tt.out {
			t.Errorf("Test: got %v, expected %v", h, tt.out)
		}
	}

}
