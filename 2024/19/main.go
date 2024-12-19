package main

import (
	"flag"
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

	part1, part2 := 0, 0
	p, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)

	var now time.Time
	var dur [2]time.Duration

	now = time.Now()
	part1, part2 = solve1(lines)
	dur[0] = time.Since(now)

	now = time.Now()
	// part2 = solve2(lines)
	dur[1] = time.Since(now)

	log.Printf("Part 1 (%v): %v\n", dur[0], part1)
	log.Printf("Part 2 (%v): %v\n", dur[1], part2)

}

func solve1(s []string) (int, int) {
	res := 0

	patterns := []string{}
	inPatterns := true
	cache := map[string]int{}

	for _, l := range s {
		if l == "" {
			inPatterns = false
			continue
		}

		if inPatterns {
			patterns = strings.Split(l, ", ")
			continue
		}

		dbg("== Line: %q", l)
		match := matchLine(l, patterns, cache)
		dbg("== %q matches: %v", l, match)
		if match > 0 {
			res++
		}
	}
	dbg("Cache: %v", cache)
	total := 0
	for _, l := range s {
		total += cache[l]
	}

	return res, total
}

func matchLine(l string, patterns []string, cache map[string]int) int {
	if len(l) == 0 {
		return 1
	}

	if v, ok := cache[l]; ok {
		return v
	}

	matches := 0
	dbg(" -> STR: %q", l)
	for _, p := range patterns {
		plen := len(p)
		if plen > len(l) {
			continue
		}
		dbg("   -> CMP %q, %q", p, l[:plen])

		if p == l[:plen] {
			dbg("   -> NEXT %q|%q", l[:plen], l[plen:])
			matches += matchLine(l[plen:], patterns, cache)
		}
	}
	cache[l] = matches

	return matches
}

func solve2(s []string) int {
	res := 0

	return res
}
