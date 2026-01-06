package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
	"strings"
	"time"

	. "github.com/agimenez/adventofcode/utils"
)

var (
	debug bool
)

func dbg(f string, v ...interface{}) {
	if debug {
		fmt.Printf(f+"\n", v...)
	}
}

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug")
}
func main() {
	flag.Parse()

	p, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)

	part1, part2, dur1, dur2 := solve(lines)
	log.Printf("Part 1 (%v): %v\n", dur1, part1)
	log.Printf("Part 2 (%v): %v\n", dur2, part2)

}

func solve(lines []string) (int, int, time.Duration, time.Duration) {
	var now time.Time
	var dur [2]time.Duration

	now = time.Now()
	part1 := solve1(lines)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 := solve2(lines)
	dur[1] = time.Since(now)

	return part1, part2, dur[0], dur[1]

}

func compileRanges(s []string) []Range {
	ranges := []Range{}
	for _, line := range s {
		ranges = append(ranges, NewRange(line))
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		return a.Cmp(b)
	})

	// Try to merge all the ranges into fully continuous ones
	// ranges is sorted, so we should be able to merge all in a single pass.
	// Start by the fist (smallest) range, and then try to merge back the subsequent ones
	merged := []Range{ranges[0]}
	for i := 1; i < len(ranges); i++ {
		cur := ranges[i]
		lastIdx := len(merged) - 1
		lastMerged := merged[lastIdx]

		if m, ok := lastMerged.MergeContiguous(cur); ok {
			merged[lastIdx] = m
		} else {
			merged = append(merged, cur)
		}
	}
	return merged
}

func solve1(s []string) int {
	res := 0

	merged := compileRanges(s)

	dbg("%v", merged)
	res = merged[0].Max() + 1

	return res
}

func solve2(s []string) int {
	res := 0

	merged := compileRanges(s)
	for _, r := range merged {
		res += r.NumValues()
	}
	res = math.MaxUint32 - res + 1
	return res
}
