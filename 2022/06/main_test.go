package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		in    string
		out4  int
		out14 int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7, 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5, 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6, 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10, 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11, 26},
	}

	for _, tt := range tests {
		start := detectStart(tt.in, 4)
		if start != tt.out4 {
			t.Errorf("%v: got %v, expected %v", tt.in, start, tt.out4)
		}
	}

}
