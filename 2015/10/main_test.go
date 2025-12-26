package main

import (
	"slices"
	"testing"

	"github.com/agimenez/adventofcode/utils"
)

func TestMain(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"1", "11"},
		{"11", "21"},
		{"21", "1211"},
		{"1211", "111221"},
		{"111221", "312211"},
	}

	for _, tt := range tests {
		in := utils.CSVToIntSlice(tt.in, "")
		out := utils.CSVToIntSlice(tt.out, "")
		l := lookAndSay(in)
		if !slices.Equal(l, out) {
			t.Errorf("Test (%v): got %v, expected %v", in, l, out)
		}
	}

}
