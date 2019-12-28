package main

import (
	"reflect"
	"testing"
)

func TestPattern(t *testing.T) {
	patterns := []struct {
		phase   int
		pattern []int
	}{
		{1, []int{0, 1, 0, -1}},
		{2, []int{0, 0, 1, 1, 0, 0, -1, -1}},
		{3, []int{0, 0, 0, 1, 1, 1, 0, 0, 0, -1, -1, -1}},
	}

	for _, n := range patterns {
		p := pattern(n.phase)

		if !reflect.DeepEqual(p, n.pattern) {
			t.Errorf("Got %#v, expected %#v", p, n.pattern)
		}

	}
}
