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

func Test_getGroups(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"{}", 1},
		{"{{{}}}", 6},
		{"{{},{}}", 5},
		{"{{{},{},{{}}}}", 16},
		{"{<a>,<a>,<a>,<a>}", 1},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got, _ := processStream(tt.s)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("getGroups(%s) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
