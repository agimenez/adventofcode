package utils

import (
	"cmp"
	"strconv"
	"strings"
)

type Range struct {
	min, max int
}

func (r Range) Min() int {
	return r.min
}

func (r Range) Max() int {
	return r.max
}

func NewRange(in string) Range {
	var r Range

	parts := strings.Split(in, "-")
	r.min, _ = strconv.Atoi(parts[0])
	r.max, _ = strconv.Atoi(parts[1])

	return r
}

func (r Range) Contains(r2 Range) bool {
	return r.min <= r2.min && r.max >= r2.max
}

func (r Range) Overlaps(r2 Range) bool {
	return r.min <= r2.min && r.max >= r2.min
}

func (r Range) Merge(r2 Range) Range {
	return Range{
		min: Min(r.min, r2.min),
		max: Max(r.max, r2.max),
	}
}

func (r Range) ContainsInt(v int) bool {
	if r.min <= v && v <= r.max {
		return true
	}

	return false
}

func (r Range) Cmp(r2 Range) int {
	return cmp.Compare(r.min, r2.min)
}

func (r Range) NumValues() int {
	return r.max - r.min + 1
}

func (r Range) ToSlice() []int {

	result := make([]int, r.max-r.min+1)
	for i := range result {
		result[i] = r.min + i
	}

	return result
}

// All returns an iterator that yields (value, true) for each integer in the range.
func (r Range) All(yield func(int, bool) bool) {
	for i := r.min; i <= r.max; i++ {
		if !yield(i, true) {
			return
		}
	}
}
