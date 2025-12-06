package utils

import (
	"strconv"
	"strings"
)

type Range struct {
	min, max int
}

func NewRange(in string) Range {
	var r Range

	parts := strings.Split(in, "-")
	r.min, _ = strconv.Atoi(parts[0])
	r.max, _ = strconv.Atoi(parts[1])

	return r
}

func (r Range) ContainsRange(r2 Range) bool {
	return r.min <= r2.min && r.max >= r2.max
}

func (r Range) Overlaps(r2 Range) bool {
	return r.min <= r2.min && r.max >= r2.min
}

func (r Range) ContainsInt(v int) bool {
	if r.min <= v && v <= r.max {
		return true
	}

	return false
}
