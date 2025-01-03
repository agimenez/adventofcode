package main

import (
	"flag"
	"io"
	"log"
	"os"
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

	part1, part2 := 0, 0
	p, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := string(p)
	dbg("lines: %#v", lines)

	var now time.Time
	var dur [2]time.Duration

	now = time.Now()
	part1 = solve1(lines)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 = solve2(lines)
	dur[1] = time.Since(now)

	log.Printf("Part 1 (%v): %v\n", dur[0], part1)
	log.Printf("Part 2 (%v): %v\n", dur[1], part2)

}

func floor(s string) int {
	cur := 0
	for _, c := range s {
		if c == '(' {
			cur++
		} else if c == ')' {
			cur--
		}
	}

	return cur
}

func solve1(s string) int {
	return floor(s)
}

func solve2(s string) int {
	cur := 0
	for i, c := range s {
		if c == '(' {
			cur++
		} else if c == ')' {
			cur--
		}
		if cur == -1 {
			return i   + 1
		}
	}

	return 0
}
