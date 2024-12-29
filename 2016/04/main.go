package main

import (
	"flag"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/agimenez/adventofcode/utils"
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

func solve1(s []string) int {
	res := 0

	for _, l := range s {
		freqs := map[rune]int{}
		parts := strings.Split(l, "-")
		name := strings.Join(parts[:len(parts)-1], "")
		dbg("Line: %v", l)
		dbg(" -> name: %v", name)
		for _, c := range name {
			freqs[c]++
		}

		charSlice := []rune{}
		for k := range freqs {
			charSlice = append(charSlice, k)
		}
		sort.Slice(charSlice, func(a, b int) bool {
			fa := freqs[charSlice[a]]
			fb := freqs[charSlice[b]]
			if fa == fb {
				return charSlice[a] < charSlice[b]
			}

			return fa > fb
		})
		dbg(" -> sorted: %v", string(charSlice))

		last := strings.Split(parts[len(parts)-1], "[")
		room := utils.ToInt(last[0])
		checksum := last[1][:len(last[1])-1]
		dbg(" -> Room: %v, check: %v", room, checksum)
		if string(charSlice[:5]) == checksum {
			res += room
		}

	}

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
