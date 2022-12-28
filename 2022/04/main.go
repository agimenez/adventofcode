package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	debug bool
)

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
	}
}

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug")
}

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

func (r Range) Contains(r2 Range) bool {
	return r.min <= r2.min && r.max >= r2.max
}

func (r Range) Overlaps(r2 Range) bool {
	return r.min <= r2.min && r.max >= r2.min
}

func ParseRanges(in string) (Range, Range) {
	ranges := strings.Split(in, ",")
	r1 := NewRange(ranges[0])
	r2 := NewRange(ranges[1])

	return r1, r2
}

func fullyContains(in string) bool {
	r1, r2 := ParseRanges(in)
	return r1.Contains(r2) || r2.Contains(r1)
}

func overlaps(in string) bool {
	r1, r2 := ParseRanges(in)

	return r1.Overlaps(r2) || r2.Overlaps(r1)
}

func main() {
	flag.Parse()

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	dbg("lines: %#v", lines)
	for i := range lines {
		if fullyContains(lines[i]) {
			part1++
		}

		if overlaps(lines[i]) {
			part2++
		}
	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
