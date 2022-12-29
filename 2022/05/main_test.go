package main

import (
	"bytes"
	"testing"
)

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

func TestNStack(t *testing.T) {
	var s Stack

	s = s.Push('a').Push('b').Push('c').Push('d')
	s, top2 := s.PopN(2)
	if len(s) != 2 {
		t.Errorf("PopN(2): Wrong size: got %d, expected %d", len(s), 2)
	}

	if !bytes.Equal(top2, []byte{'c', 'd'}) {
		t.Errorf("PopN(2): Wrong elems: got %v, expected %v", top2, []byte{'c', 'd'})
	}

	s = s.PushN([]byte{'e', 'f'})
	if !bytes.Equal(s, []byte{'a', 'b', 'e', 'f'}) {
		t.Errorf("PushN: got %v, expected %v", s, []byte{'a', 'b', 'e', 'f'})
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
