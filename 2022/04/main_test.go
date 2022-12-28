package main

import (
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		in       string
		r1       Range
		r2       Range
		contains bool
		overlap  bool
	}{
		{"2-4,6-8", Range{2, 4}, Range{6, 8}, false, false},
		{"2-3,4-5", Range{2, 3}, Range{4, 5}, false, false},
		{"5-7,7-9", Range{5, 7}, Range{7, 9}, false, true},
		{"2-8,3-7", Range{2, 8}, Range{3, 7}, true, true},
		{"6-6,4-6", Range{6, 6}, Range{4, 6}, true, true},
		{"2-6,4-8", Range{2, 6}, Range{4, 8}, false, true},
		{"96-97,26-64", Range{96, 97}, Range{26, 64}, false, false},
	}

	for _, tt := range tests {
		parts := strings.Split(tt.in, ",")
		r1 := NewRange(parts[0])
		if r1 != tt.r1 {
			t.Errorf("Range1 %s: got %v, expected %v", parts[0], r1, tt.r1)
		}

		r2 := NewRange(parts[1])
		if r2 != tt.r2 {
			t.Errorf("Range2 %s: got %v, expected %v", parts[1], r2, tt.r2)
		}

		c := fullyContains(tt.in)
		if c != tt.contains {
			t.Errorf("fullyContains %q: got %v, expected %v", tt.in, c, tt.contains)
		}

		o := overlaps(tt.in)
		if o != tt.overlap {
			t.Errorf("overlaps %q: got %v, expected %v", tt.in, o, tt.overlap)
		}
	}

}
