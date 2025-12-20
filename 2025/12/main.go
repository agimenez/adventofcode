package main

import (
	"flag"
	. "github.com/agimenez/adventofcode/utils"
	"io"
	"log"
	"os"
	"strings"
	"time"
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
	shape_area := 9 // yeah, this sucks, but it is what it is
	for _, line := range s {
		region := strings.Split(line, ": ")
		size := strings.Split(region[0], "x")
		// jump shapes section, we know it's 3x3
		if len(size) < 2 {
			continue
		}
		dbg("Got region: %v", line)

		region_area := ToInt(size[0]) * ToInt(size[1])
		required_counts := strings.Fields(region[1])
		total_required_count := 0
		for _, count := range required_counts {
			total_required_count += ToInt(count)
		}
		dbg(" >> Total required count: %v, count area: %v, region area: %v", total_required_count, total_required_count*shape_area, region_area)
		if total_required_count*shape_area <= region_area {
			res++
		}

	}

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
