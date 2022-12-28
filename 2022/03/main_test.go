package main

import "testing"

func TestDups(t *testing.T) {
	tests := []struct {
		in   string
		dup  rune
		prio int
	}{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", 'p', 16},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", 'L', 38},
		{"PmmdzqPrVvPwwTWBwg", 'P', 42},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", 'v', 22},
		{"ttgJtRGJQctTZtZT", 't', 20},
		{"CrZsJsPPZsGzwwsLwLmpwMDw", 's', 19},
	}

	for _, tt := range tests {
		r := NewRuckSack(tt.in)
		dup := r.getDuplicate()
		if dup != tt.dup {
			t.Errorf("Dup: got %v, expected %v", dup, tt.dup)
		}

		prio := priority(dup)
		if prio != tt.prio {
			t.Errorf("Prio (%q): got %v, expected %v", dup, prio, tt.prio)
		}
	}

}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		in   []string
		dup  rune
		prio int
	}{
		{[]string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"}, 'r', 18},
		{[]string{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}, 'Z', 52},
	}

	for _, tt := range tests {
		c := findCommon(tt.in)
		if c != tt.dup {
			t.Errorf("Common: got %v, expected %v", c, tt.dup)
		}

		prio := priority(c)
		if prio != tt.prio {
			t.Errorf("Prio: got %v, expected %v", prio, tt.prio)
		}

	}
}
