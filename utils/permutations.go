package utils

import "iter"

// Permutations returns an iterator over all permutations of the input slice.
// Uses Heap's algorithm for efficient generation.
// Each yielded slice is a new copy, safe to store or modify.
func Permutations[T any](s []T) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		if len(s) == 0 {
			return
		}

		// Work on a copy to avoid modifying the original
		arr := make([]T, len(s))
		copy(arr, s)

		// Heap's algorithm using an explicit stack instead of recursion
		n := len(arr)
		c := make([]int, n) // control array

		// Yield the first permutation
		tmp := make([]T, n)
		copy(tmp, arr)
		if !yield(tmp) {
			return
		}

		i := 0
		for i < n {
			if c[i] < i {
				if i%2 == 0 {
					arr[0], arr[i] = arr[i], arr[0]
				} else {
					arr[c[i]], arr[i] = arr[i], arr[c[i]]
				}

				// Yield current permutation
				tmp := make([]T, n)
				copy(tmp, arr)
				if !yield(tmp) {
					return
				}

				c[i]++
				i = 0
			} else {
				c[i] = 0
				i++
			}
		}
	}
}

// PermutationsIndexed returns an iterator over all permutations with their index.
// Index starts at 0 and increments for each permutation.
func PermutationsIndexed[T any](s []T) iter.Seq2[int, []T] {
	return func(yield func(int, []T) bool) {
		idx := 0
		for perm := range Permutations(s) {
			if !yield(idx, perm) {
				return
			}
			idx++
		}
	}
}

// PermutationsInPlace returns an iterator that yields permutations without copying.
// WARNING: The yielded slice is reused between iterations - do not store it directly.
// Use this for better performance when you only need to read each permutation once.
func PermutationsInPlace[T any](s []T) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		if len(s) == 0 {
			return
		}

		// Work on a copy to avoid modifying the original
		arr := make([]T, len(s))
		copy(arr, s)

		n := len(arr)
		c := make([]int, n)

		// Yield the first permutation
		if !yield(arr) {
			return
		}

		i := 0
		for i < n {
			if c[i] < i {
				if i%2 == 0 {
					arr[0], arr[i] = arr[i], arr[0]
				} else {
					arr[c[i]], arr[i] = arr[i], arr[c[i]]
				}

				if !yield(arr) {
					return
				}

				c[i]++
				i = 0
			} else {
				c[i] = 0
				i++
			}
		}
	}
}

// CountPermutations returns the number of permutations for a slice of length n.
// This is n! (factorial).
func CountPermutations(n int) int {
	if n <= 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// CollectPermutations collects all permutations into a slice.
// Convenience function when you need all permutations in memory.
func CollectPermutations[T any](s []T) [][]T {
	result := make([][]T, 0, CountPermutations(len(s)))
	for perm := range Permutations(s) {
		result = append(result, perm)
	}
	return result
}
