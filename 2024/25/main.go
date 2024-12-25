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

	var keys, locks [][5]int

	for i := 0; i < len(lines); i += 8 {
		var cur [5]int

		dbg("NEW line %v: %s", i, lines[i])
		for j := 0; j < 7; j++ {
			dbg(" -> j=%v, j=%s", j, lines[i+j])
			for col, ch := range lines[i+j] {
				if ch == '#' {
					cur[col]++
				}
			}
		}

		if lines[i][0] == '#' {
			locks = append(locks, cur)
		} else {
			keys = append(keys, cur)
		}
	}

	dbg("keys: %v", keys)
	dbg("locks: %v", locks)

	var now time.Time
	var dur [2]time.Duration

	now = time.Now()
	part1 = solve1(keys, locks)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 = solve2(lines)
	dur[1] = time.Since(now)

	log.Printf("Part 1 (%v): %v\n", dur[0], part1)
	log.Printf("Part 2 (%v): %v\n", dur[1], part2)

}

func solve1(keys, locks [][5]int) int {
	res := 0

	for _, lock := range locks {
		dbg("Lock %v", lock)
		for _, key := range keys {
			dbg(" -> Key %v", key)
			fits := true
			for pin := range key {
				if lock[pin]+key[pin] > 7 {
					fits = false
					break
				}
			}
			if fits {
				res++
			}
		}
	}

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
