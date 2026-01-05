package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
	// . "github.com/agimenez/adventofcode/utils"
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

func NextRow(s string) string {
	var b strings.Builder

	for i := range s {
		var l, r byte
		if i == 0 {
			l = '.'
		} else {
			l = s[i-1]
		}

		if i == len(s)-1 {
			r = '.'
		} else {
			r = s[i+1]
		}
		c := s[i]

		// This is some sort of XOR, and the center tile does't
		// do anything really, but let's stick to the rules explicitly
		out := '.'
		if l == '^' && c == '^' && r != '^' {
			out = '^'
		}

		if c == '^' && r == '^' && l != '^' {
			out = '^'
		}

		if l == '^' && c != '^' && r != '^' {
			out = '^'
		}

		if l != '^' && c != '^' && r == '^' {
			out = '^'
		}

		b.WriteRune(out)
	}

	return b.String()
}

func CountSafe(s string, rows int) int {
	out := 0
	for r := range rows {
		dbg("Row %d: %s", r+1, s)
		out += strings.Count(s, ".")
		s = NextRow(s)
	}

	return out
}

func solve1(s []string) int {
	res := 0

	res = CountSafe(s[0], 40)

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
