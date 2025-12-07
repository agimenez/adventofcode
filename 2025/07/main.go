package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
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

func str2grid(lines []string) [][]rune {
	result := make([][]rune, len(lines))
	for i, line := range lines {
		result[i] = []rune(line)
	}
	return result
}

func printGrid(grid [][]rune) {
	if !debug {
		return
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}
	println()
}

func solve1(s []string) int {
	res := 0

	grid := str2grid(s)
	splitters := map[Point]bool{}
	beams := map[Point]bool{}
	for y, line := range s {
		for x, c := range line {
			p := Point{x, y}
			if c == 'S' {
				beams[p] = true
				continue
			}

			if c == '^' {
				splitters[p] = true
			}
		}
	}
	dbg("Initial beam loc: %v", beams)
	dbg("Splitters: %v", splitters)

	done := false
	for !done {
		printGrid(grid)
		done = true
		nextBeams := map[Point]bool{}

		// dbg(">> Current tracked beams: %v", beams)
		for beam := range beams {
			newPos := beam.Down()
			ch, in := GetChInPoint(s, newPos)
			dbg(">>>> Checking beam %v (%c)", beam, ch)

			// Not outside of the grid, so keep checking
			if in {
				if ch == '^' {
					// Found a splitter, so generete new beams on the sides of it
					// If there's already a beam there, we don't want to count it twice
					if _, ok := beams[newPos.Left()]; !ok {
						nextBeams[newPos.Left()] = true
						dbg("  >>> New split for %v", newPos.Left())
					}
					if _, ok := beams[newPos.Right()]; !ok {
						nextBeams[newPos.Right()] = true
						dbg("  >>> New split for %v", newPos.Right())
					}
					done = false
					res++

					grid[newPos.Left().Y][newPos.Left().X] = '|'
					grid[newPos.Right().Y][newPos.Right().X] = '|'
				} else {
					// Nothingness, so keep going down
					nextBeams[newPos] = true
					done = false
					grid[newPos.Y][newPos.X] = '|'

				}
			}
			// dbg(">>>> Current nextBeams: %v, splits = %v", nextBeams, res)
			dbg(">>>> Current splits = %v", res)
		}
		beams = nextBeams
	}

	return res
}

func solve2(s []string) int {
	res := 0

	// Go line by line as we fall down. This is the "previous generation" of timelines
	// it contains all the timlines that exist in that column of the grid.
	// Whenever we encounter a splitter, will split also the counts to the 2 adjacent slots
	prevBeams := make([]int, len(s[0]))
	for _, line := range s {
		currentBeams := make([]int, len(line))
		for x, c := range line {
			if c == 'S' {
				currentBeams[x] = 1
			} else if c == '^' {
				currentBeams[x-1] += prevBeams[x]
				currentBeams[x+1] += prevBeams[x]
			} else {
				currentBeams[x] += prevBeams[x]
			}
		}

		prevBeams = currentBeams
	}

	for _, n := range prevBeams {
		res += n
	}

	return res
}
