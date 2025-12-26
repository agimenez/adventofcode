package utils

import (
	"slices"
	"testing"
)

func TestPermutations(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int // expected number of permutations
	}{
		{"empty", []int{}, 0},
		{"single", []int{1}, 1},
		{"two", []int{1, 2}, 2},
		{"three", []int{1, 2, 3}, 6},
		{"four", []int{1, 2, 3, 4}, 24},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count := 0
			for range Permutations(tt.input) {
				count++
			}
			if count != tt.expected {
				t.Errorf("Permutations(%v) yielded %d permutations, want %d", tt.input, count, tt.expected)
			}
		})
	}
}

func TestPermutationsContent(t *testing.T) {
	input := []int{1, 2, 3}
	expected := [][]int{
		{1, 2, 3},
		{2, 1, 3},
		{3, 1, 2},
		{1, 3, 2},
		{2, 3, 1},
		{3, 2, 1},
	}

	var result [][]int
	for perm := range Permutations(input) {
		result = append(result, perm)
	}

	if len(result) != len(expected) {
		t.Fatalf("got %d permutations, want %d", len(result), len(expected))
	}

	// Check all expected permutations are present
	for _, exp := range expected {
		found := false
		for _, res := range result {
			if slices.Equal(res, exp) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("missing permutation %v", exp)
		}
	}
}

func TestPermutationsStrings(t *testing.T) {
	input := []string{"a", "b", "c"}
	count := 0
	for range Permutations(input) {
		count++
	}
	if count != 6 {
		t.Errorf("got %d permutations, want 6", count)
	}
}

func TestPermutationsIndexed(t *testing.T) {
	input := []int{1, 2, 3}
	indices := []int{}
	for idx, _ := range PermutationsIndexed(input) {
		indices = append(indices, idx)
	}

	expected := []int{0, 1, 2, 3, 4, 5}
	if !slices.Equal(indices, expected) {
		t.Errorf("indices = %v, want %v", indices, expected)
	}
}

func TestPermutationsInPlace(t *testing.T) {
	input := []int{1, 2, 3}
	count := 0
	for range PermutationsInPlace(input) {
		count++
	}
	if count != 6 {
		t.Errorf("got %d permutations, want 6", count)
	}
}

func TestPermutationsEarlyExit(t *testing.T) {
	input := []int{1, 2, 3, 4}
	count := 0
	for range Permutations(input) {
		count++
		if count == 5 {
			break
		}
	}
	if count != 5 {
		t.Errorf("early exit: got %d iterations, want 5", count)
	}
}

func TestCountPermutations(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
	}

	for _, tt := range tests {
		if got := CountPermutations(tt.n); got != tt.expected {
			t.Errorf("CountPermutations(%d) = %d, want %d", tt.n, got, tt.expected)
		}
	}
}

func TestCollectPermutations(t *testing.T) {
	input := []int{1, 2, 3}
	result := CollectPermutations(input)
	if len(result) != 6 {
		t.Errorf("CollectPermutations(%v) returned %d permutations, want 6", input, len(result))
	}
}

func TestPermutationsDoesNotModifyInput(t *testing.T) {
	input := []int{1, 2, 3}
	original := slices.Clone(input)

	for range Permutations(input) {
		// consume all
	}

	if !slices.Equal(input, original) {
		t.Errorf("input was modified: got %v, want %v", input, original)
	}
}

func BenchmarkPermutations(b *testing.B) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range Permutations(input) {
		}
	}
}

func BenchmarkPermutationsInPlace(b *testing.B) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range PermutationsInPlace(input) {
		}
	}
}
