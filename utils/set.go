package utils

import "iter"
import "maps"

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return Set[T]{}
}

func (s Set[T]) Len() int {
	return len(s)
}

// Add a new element of type T to the set.
func (s Set[T]) Add(v T) bool {
	l := len(s)

	s[v] = struct{}{}

	return l != len(s)
}

func NewSetFromMapKeys[T comparable, V any](m map[T]V) Set[T] {
	s := make(Set[T], len(m))

	for k := range m {
		s.Add(k)
	}

	return s
}

func (s Set[T]) Remove(v T) {
	delete(s, v)
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]

	return ok
}

// Difference returns a set with elements from `s` that are not in `other`.
func (s Set[T]) Difference(other Set[T]) Set[T] {
	diff := Set[T]{}

	for v := range s {
		if !other.Contains(v) {
			diff.Add(v)
		}
	}

	return diff
}

// Returns an iterator with all the elements
func (s Set[T]) All() iter.Seq[T] {
	return maps.Keys(s)
}
