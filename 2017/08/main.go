package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"maps"
	"os"
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

func cond(s string, reg map[string]int) bool {
	parts := strings.Fields(s)
	r := reg[parts[0]]
	v := ToInt(parts[2])

	switch parts[1] {
	case ">":
		return r > v
	case "<":
		return r < v
	case ">=":
		return r >= v
	case "==":
		return r == v
	case "<=":
		return r <= v
	case "!=":
		return r != v

	default:
		panic("Unknown comparator: " + parts[1])
	}

}

func exec(s string, reg map[string]int) int {
	parts := strings.Fields(s)
	r := parts[0]
	v := ToInt(parts[2])

	switch parts[1] {
	case "inc":
		reg[r] += v

	case "dec":
		reg[r] -= v

	default:
		panic("Unknown op: " + parts[1])

	}

	return reg[r]
}

func solve1(s []string) int {
	res := 0

	reg := map[string]int{}
	for _, line := range s {
		parts := strings.Split(line, " if ")
		if cond(parts[1], reg) {
			exec(parts[0], reg)
		}

	}
	dbg("%v", reg)

	for v := range maps.Values(reg) {
		res = Max(res, v)
	}

	dbg("")
	return res
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")
	reg := map[string]int{}
	for _, line := range s {
		parts := strings.Split(line, " if ")
		if cond(parts[1], reg) {
			res = Max(res, exec(parts[0], reg))
		}

	}
	dbg("%v", reg)

	return res
}
