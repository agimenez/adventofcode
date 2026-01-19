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

func TestCountSteps(t *testing.T) {
	tests := []struct {
		path string
		want int
	}{
		{"ne,ne,ne", 3},
		{"ne,ne,sw,sw", 0},
		{"ne,ne,s,s", 2},
		{"se,sw,se,sw,sw", 3},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			got := CountSteps(tt.path)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("CountSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
