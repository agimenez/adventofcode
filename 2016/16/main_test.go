package main

import "testing"

func TestDragonCurve(t *testing.T) {
	tests := []struct {
		in  string
		len int

		out string
	}{
		{"10000", 20, "10000011110010000111"},
	}

	for _, tt := range tests {
		l := DragonCurve(tt.in, tt.len)
		if l != tt.out {
			t.Errorf("Test (%v): got %v, expected %v", tt.in, l, tt.out)
		}
	}

}

func TestCheckSum(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"110010110100", "100"},
		{"10000011110010000111", "01100"},
	}

	for _, tt := range tests {
		l := CheckSum(tt.in)
		if l != tt.out {
			t.Errorf("Test (%v): got %v, expected %v", tt.in, l, tt.out)
		}
	}
}

func TestDragonStep(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want string
	}{
		{"1", "100"},
		{"0", "001"},
		{"11111", "11111000000"},
		{"111100001010", "1111000010100101011110000"},
		{"10000", "10000011110"},
		{"10000011110", "10000011110010000111110"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DragonStep(tt.name)
			if got != tt.want {
				t.Errorf("DragonCurve() = %v, want %v", got, tt.want)
			}
		})
	}
}
