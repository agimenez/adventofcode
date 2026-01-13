package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa", false},
		{"aa bb cc dd aaaa", true},
	}

	for _, tt := range tests {
		l := valid(tt.in)
		if l != tt.out {
			t.Errorf("Test (%v): got %v, expected %v", tt.in, l, tt.out)
		}
	}

}

func Test_noAnagrams(t *testing.T) {
	tests := []struct {
		pass string
		want bool
	}{
		{"abcde fghij", true},
		{"abcde xyz ecdab", false},
		{"a ab abc abd abf abj", true},
		{"iiii oiii ooii oooi oooo", true},
		{"oiii ioii iioi iiio", false},
	}
	for _, tt := range tests {
		t.Run(tt.pass, func(t *testing.T) {
			got := noAnagrams(tt.pass)
			if got != tt.want {
				t.Errorf("noAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
