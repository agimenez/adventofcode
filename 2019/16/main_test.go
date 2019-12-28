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

func TestFFT(t *testing.T) {
	ffts := []struct {
		phases int
		input  string
		output string
	}{
		{4, "12345678", "01029498"},
		{100, "80871224585914546619083218645595", "24176176"},
		{100, "19617804207202209144916044189917", "73745418"},
		{100, "69317163492948606335995924319873", "52432133"},
	}

	for _, fft := range ffts {
		res := FFT(fft.input, fft.phases)

		if res[:8] != fft.output {
			t.Errorf("Got %v, expected %v", res, fft.output)
		}
	}
}
