package main

import (
	"flag"
	"io"
	"log"
	"os"
	"sort"
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

func solve1(s []string) int {
	res := 0

	for _, triangle := range s {
		parts := strings.Fields(triangle)
		sides := []int{
			ToInt(parts[0]),
			ToInt(parts[1]),
			ToInt(parts[2]),
		}
		sort.Ints(sides)
		if sides[0]+sides[1] > sides[2] {
			res++
		}

	}

	return res
}

func solve2(s []string) int {
	res := 0

	triangles := [3][3]int{}
	for r, line := range s {
		parts := strings.Fields(line)

		idx := r % 3
		for triangle, side := range parts {
			triangles[triangle][idx] = ToInt(side)
		}
		dbg("Triangles: %+v", triangles)

		if (r+1)%3 == 0 {
			dbg(" -> CALC")
			for _, triangle := range triangles {
				sort.Ints(triangle[:])
				dbg(" -> %v", triangle)
				if triangle[0]+triangle[1] > triangle[2] {
					res++
				}
			}
		}

	}

	return res
}
