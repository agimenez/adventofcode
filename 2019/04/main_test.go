package main

import "testing"

func TestCompliance(t *testing.T) {
	numbers := []struct {
		c  int
		ok bool
	}{
		{112233, true},
		{123444, false},
		{111122, true},
	}

	for _, n := range numbers {
		c := checkCompliant(n.c)

		if n.ok == true && c == false {
			t.Errorf("%d expected to be compliant, detected not", n.c)
		}

		if n.ok == false && c == true {
			t.Errorf("%d expected non compliant, detected it is", n.c)
		}
	}
}
