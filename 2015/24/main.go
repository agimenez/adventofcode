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

func BalancePackages(nums []int, groups int) int {
	totalSum := Reduce(nums, 0, ReduceSum)
	sumByGroup := totalSum / groups
	maxElems := len(nums) / groups
	dbg("Total Sum: %v, sum by group: %v, total elems: %v, max elems in G1: %v", totalSum, sumByGroup, len(nums), maxElems)

	slices.Sort(nums)
	slices.Reverse(nums)

	minEntanglement := math.MaxInt
	minLen := math.MaxInt
	found := false
	for i := 1; i <= maxElems && !found; i++ {
		for group1 := range Combinations(nums, i) {
			sum := Reduce(group1, 0, ReduceSum)
			dbg("Testing %v (sum: %v, required: %v)", group1, sum, sumByGroup)
			if sum != sumByGroup {
				dbg(" >> Sum mismatch!")
				continue
			}

			if len(group1) > minLen {
				dbg(" >> We have a shorter combination!")
				continue
			}

			entanglement := Reduce(group1, 1, ReduceMult)
			minEntanglement = Min(minEntanglement, entanglement)
			minLen = len(group1)
			dbg("NEW! %v -> %v (%v)", group1, entanglement, minEntanglement)
			found = true
		}

	}

	return minEntanglement
}

func solve1(s []string) int {
	nums := LinesToIntSlice(s)

	return BalancePackages(nums, 3)
}

func solve2(s []string) int {
	nums := LinesToIntSlice(s)

	return BalancePackages(nums, 4)
}
