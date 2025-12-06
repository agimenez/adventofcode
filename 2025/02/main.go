package main

import (
	"flag"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	. "github.com/agimenez/adventofcode/utils"
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

func isValid(n int) bool {
	res := true

	str := strconv.Itoa(n)
	strlen := len(str)
	if strlen%2 == 0 {
		part1 := str[:strlen/2]
		part2 := str[strlen/2:]

		if part1 == part2 {
			return false
		}
	}

	return res
}

func isValid2(n int) bool {
	valid := true
	id := strconv.Itoa(n)

	// if len(id) == 1 {
	// 	return true
	// }

	characters := strings.Split(id, "")

	for chunkSize := 1; chunkSize <= (len(characters)/2) && valid; chunkSize++ {
		chunks := slices.Collect(slices.Chunk(characters, chunkSize))

		// We only need to know how many chunks are different to the first
		// Anything more than 0 means the ID is valid
		differentToFirst := 0

		for c := 1; c < len(chunks); c++ {
			if !slices.Equal(chunks[c], chunks[0]) {
				differentToFirst++
			}
		}

		if differentToFirst == 0 {
			valid = false
		}
	}

	return valid
}

func solve1(s []string) int {
	res := 0
	ranges := strings.Split(s[0], ",")
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		min := ToInt(parts[0])
		max := ToInt(parts[1])

		for ; min <= max; min++ {
			if !isValid(min) {
				res += min
			}
		}
	}

	return res
}

func solve2(s []string) int {
	res := 0

	ranges := strings.Split(s[0], ",")
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		min := ToInt(parts[0])
		max := ToInt(parts[1])

		for ; min <= max; min++ {
			if !isValid2(min) {
				res += min
			}
		}
	}

	return res
}
