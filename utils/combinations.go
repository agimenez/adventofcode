package utils

import "iter"

// NOTE: AI generated
//
// Combinations returns an iterator over all combinations of k elements from slice s.
// Combinations are generated in lexicographic order.
// Each yielded slice is a new copy, safe to store or modify.
//
// If k > len(s) or k < 0, no combinations are yielded.
// If k == 0, yields one empty slice (there's exactly one way to choose nothing).
func Combinations[T any](s []T, k int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		n := len(s)

		// Edge cases
		if k < 0 || k > n {
			return
		}
		if k == 0 {
			yield([]T{})
			return
		}

		// Initialize indices to [0, 1, 2, ..., k-1]
		indices := make([]int, k)
		for i := range indices {
			indices[i] = i
		}

		for {
			// Build and yield current combination
			combo := make([]T, k)
			for i, idx := range indices {
				combo[i] = s[idx]
			}
			if !yield(combo) {
				return
			}

			// Find rightmost index that can be incremented
			// indices[i] can go up to n - k + i
			i := k - 1
			for i >= 0 && indices[i] == n-k+i {
				i--
			}

			// All combinations generated
			if i < 0 {
				return
			}

			// Increment this index and reset all indices to its right
			indices[i]++
			for j := i + 1; j < k; j++ {
				indices[j] = indices[j-1] + 1
			}
		}
	}
}

// CombinationsIndexed returns an iterator over all combinations with their index.
// Index starts at 0 and increments for each combination.
func CombinationsIndexed[T any](s []T, k int) iter.Seq2[int, []T] {
	return func(yield func(int, []T) bool) {
		idx := 0
		for combo := range Combinations(s, k) {
			if !yield(idx, combo) {
				return
			}
			idx++
		}
	}
}

// CombinationsInPlace returns an iterator that yields combinations without copying.
// WARNING: The yielded slice is reused between iterations - do not store it directly.
// Use this for better performance when you only need to read each combination once.
func CombinationsInPlace[T any](s []T, k int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		n := len(s)

		if k < 0 || k > n {
			return
		}
		if k == 0 {
			yield([]T{})
			return
		}

		// Initialize indices to [0, 1, 2, ..., k-1]
		indices := make([]int, k)
		for i := range indices {
			indices[i] = i
		}

		// Reusable combination slice
		combo := make([]T, k)

		for {
			// Build current combination
			for i, idx := range indices {
				combo[i] = s[idx]
			}
			if !yield(combo) {
				return
			}

			// Find rightmost index that can be incremented
			i := k - 1
			for i >= 0 && indices[i] == n-k+i {
				i--
			}

			if i < 0 {
				return
			}

			indices[i]++
			for j := i + 1; j < k; j++ {
				indices[j] = indices[j-1] + 1
			}
		}
	}
}

// CombinationsIndices returns an iterator over index combinations.
// Useful when you need the indices rather than the elements themselves.
// Each yielded slice is a new copy.
func CombinationsIndices(n, k int) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		if k < 0 || k > n {
			return
		}
		if k == 0 {
			yield([]int{})
			return
		}

		indices := make([]int, k)
		for i := range indices {
			indices[i] = i
		}

		for {
			// Copy and yield
			combo := make([]int, k)
			copy(combo, indices)
			if !yield(combo) {
				return
			}

			i := k - 1
			for i >= 0 && indices[i] == n-k+i {
				i--
			}

			if i < 0 {
				return
			}

			indices[i]++
			for j := i + 1; j < k; j++ {
				indices[j] = indices[j-1] + 1
			}
		}
	}
}

// CountCombinations returns the number of combinations C(n, k) = n! / (k! * (n-k)!).
// Returns 0 for invalid inputs (k < 0 or k > n).
func CountCombinations(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}

	// Use the smaller of k and n-k for efficiency
	if k > n-k {
		k = n - k
	}

	// Calculate C(n,k) = n * (n-1) * ... * (n-k+1) / k!
	// Do multiplication and division in steps to avoid overflow
	result := 1
	for i := 0; i < k; i++ {
		result *= (n - i)
		result /= (i + 1)
	}

	return result
}

// CollectCombinations collects all combinations into a slice.
// Convenience function when you need all combinations in memory.
func CollectCombinations[T any](s []T, k int) [][]T {
	result := make([][]T, 0, CountCombinations(len(s), k))
	for combo := range Combinations(s, k) {
		result = append(result, combo)
	}
	return result
}

// CombinationsAll returns an iterator over all combinations of all lengths (0 to n).
// This generates the power set minus ordering.
func CombinationsAll[T any](s []T) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		for k := 0; k <= len(s); k++ {
			for combo := range Combinations(s, k) {
				if !yield(combo) {
					return
				}
			}
		}
	}
}
