package main

import "testing"

func TestStack(t *testing.T) {
	var s Stack

	s = s.Push('a')
	if top := s.Top(); top != 'a' {
		t.Errorf("Top(): got %v, expected %v", top, 'a')
	}
	s, top := s.Pop()
	if top != 'a' {
		t.Errorf("Pop(): got %v, expected %v", top, 'a')
	}

	if len(s) != 0 {
		t.Errorf("Pop(): got length %v, expected 0", len(s))
	}

	s = s.Push('a').Push('b').Push('c')
	s, top = s.Pop()
	if top != 'c' {
		t.Errorf("Pop(): got %v, expected 'c'", top)
	}
	s, top = s.Pop()
	if top != 'b' {
		t.Errorf("Pop(): got %v, expected 'b'", top)
	}

	s = s.Insert('d')
	s, top = s.Pop()
	if top != 'a' {
		t.Errorf("Pop(): got %v, expected 'a'", top)
	}

	top = s.Top()
	if top != 'd' {
		t.Errorf("Insert(): got %v, expected 'd'", top)
	}

}

func TestMain(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", 16},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", 38},
		{"PmmdzqPrVvPwwTWBwg", 42},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", 22},
		{"ttgJtRGJQctTZtZT", 20},
		{"CrZsJsPPZsGzwwsLwLmpwMDw", 19},
	}

	for _, tt := range tests {
		out := tt.out
		if out != tt.out {
			t.Errorf("Test: got %v, expected %v", out, tt.out)
		}
	}

}
