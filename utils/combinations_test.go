package utils

import (
	"slices"
	"testing"
)

// NOTE: AI generated
func TestCombinations(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		k        int
		expected int // expected number of combinations
	}{
		{"empty k=0", []int{}, 0, 1},
		{"empty k=1", []int{}, 1, 0},
		{"single k=0", []int{1}, 0, 1},
		{"single k=1", []int{1}, 1, 1},
		{"single k=2", []int{1}, 2, 0},
		{"three k=0", []int{1, 2, 3}, 0, 1},
		{"three k=1", []int{1, 2, 3}, 1, 3},
		{"three k=2", []int{1, 2, 3}, 2, 3},
		{"three k=3", []int{1, 2, 3}, 3, 1},
		{"five k=2", []int{1, 2, 3, 4, 5}, 2, 10},
		{"five k=3", []int{1, 2, 3, 4, 5}, 3, 10},
		{"negative k", []int{1, 2, 3}, -1, 0},
		{"k > n", []int{1, 2, 3}, 5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count := 0
			for range Combinations(tt.input, tt.k) {
				count++
			}
			if count != tt.expected {
				t.Errorf("Combinations(%v, %d) yielded %d combinations, want %d",
					tt.input, tt.k, count, tt.expected)
			}
		})
	}
}

func TestCombinationsContent(t *testing.T) {
	input := []int{1, 2, 3, 4}
	k := 2
	expected := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 4},
		{3, 4},
	}

	var result [][]int
	for combo := range Combinations(input, k) {
		result = append(result, combo)
	}

	if len(result) != len(expected) {
		t.Fatalf("got %d combinations, want %d", len(result), len(expected))
	}

	// Combinations should be in lexicographic order
	for i, exp := range expected {
		if !slices.Equal(result[i], exp) {
			t.Errorf("combination %d: got %v, want %v", i, result[i], exp)
		}
	}
}

func TestCombinationsStrings(t *testing.T) {
	input := []string{"a", "b", "c", "d"}
	expected := [][]string{
		{"a", "b"},
		{"a", "c"},
		{"a", "d"},
		{"b", "c"},
		{"b", "d"},
		{"c", "d"},
	}

	var result [][]string
	for combo := range Combinations(input, 2) {
		result = append(result, combo)
	}

	for i, exp := range expected {
		if !slices.Equal(result[i], exp) {
			t.Errorf("combination %d: got %v, want %v", i, result[i], exp)
		}
	}
}

func TestCombinationsK0(t *testing.T) {
	input := []int{1, 2, 3}
	var result [][]int
	for combo := range Combinations(input, 0) {
		result = append(result, combo)
	}

	if len(result) != 1 {
		t.Fatalf("got %d combinations, want 1", len(result))
	}
	if len(result[0]) != 0 {
		t.Errorf("got %v, want empty slice", result[0])
	}
}

func TestCombinationsIndexed(t *testing.T) {
	input := []int{1, 2, 3, 4}
	k := 2
	indices := []int{}
	for idx, _ := range CombinationsIndexed(input, k) {
		indices = append(indices, idx)
	}

	expected := []int{0, 1, 2, 3, 4, 5}
	if !slices.Equal(indices, expected) {
		t.Errorf("indices = %v, want %v", indices, expected)
	}
}

func TestCombinationsInPlace(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	k := 3
	count := 0
	for range CombinationsInPlace(input, k) {
		count++
	}
	if count != 10 {
		t.Errorf("got %d combinations, want 10", count)
	}
}

func TestCombinationsIndices(t *testing.T) {
	expected := [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 2},
		{1, 3},
		{2, 3},
	}

	var result [][]int
	for indices := range CombinationsIndices(4, 2) {
		result = append(result, indices)
	}

	if len(result) != len(expected) {
		t.Fatalf("got %d combinations, want %d", len(result), len(expected))
	}

	for i, exp := range expected {
		if !slices.Equal(result[i], exp) {
			t.Errorf("indices %d: got %v, want %v", i, result[i], exp)
		}
	}
}

func TestCombinationsEarlyExit(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	k := 3
	count := 0
	for range Combinations(input, k) {
		count++
		if count == 5 {
			break
		}
	}
	if count != 5 {
		t.Errorf("early exit: got %d iterations, want 5", count)
	}
}

func TestCountCombinations(t *testing.T) {
	tests := []struct {
		n, k     int
		expected int
	}{
		{0, 0, 1},
		{5, 0, 1},
		{5, 5, 1},
		{5, 1, 5},
		{5, 2, 10},
		{5, 3, 10},
		{5, 4, 5},
		{10, 3, 120},
		{10, 5, 252},
		{5, -1, 0},
		{5, 6, 0},
	}

	for _, tt := range tests {
		if got := CountCombinations(tt.n, tt.k); got != tt.expected {
			t.Errorf("CountCombinations(%d, %d) = %d, want %d", tt.n, tt.k, got, tt.expected)
		}
	}
}

func TestCollectCombinations(t *testing.T) {
	input := []int{1, 2, 3, 4}
	result := CollectCombinations(input, 2)
	if len(result) != 6 {
		t.Errorf("CollectCombinations(%v, 2) returned %d combinations, want 6", input, len(result))
	}
}

func TestCombinationsAll(t *testing.T) {
	input := []int{1, 2, 3}
	// Power set has 2^n elements = 8
	count := 0
	for range CombinationsAll(input) {
		count++
	}
	if count != 8 {
		t.Errorf("CombinationsAll(%v) yielded %d, want 8", input, count)
	}
}

func TestCombinationsDoesNotModifyInput(t *testing.T) {
	input := []int{1, 2, 3, 4}
	original := slices.Clone(input)

	for range Combinations(input, 2) {
		// consume all
	}

	if !slices.Equal(input, original) {
		t.Errorf("input was modified: got %v, want %v", input, original)
	}
}

func BenchmarkCombinations(b *testing.B) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k := 5 // C(10,5) = 252
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range Combinations(input, k) {
		}
	}
}

func BenchmarkCombinationsInPlace(b *testing.B) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k := 5 // C(10,5) = 252
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range CombinationsInPlace(input, k) {
		}
	}
}
