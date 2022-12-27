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
