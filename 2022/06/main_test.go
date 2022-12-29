package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		in   string
		out4 int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, tt := range tests {
		start := detectStart(tt.in, 4)
		if start != tt.out4 {
			t.Errorf("%v: got %v, expected %v", tt.in, start, tt.out4)
		}
	}

}
