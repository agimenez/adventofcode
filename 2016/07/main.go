package main

import (
	"flag"
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

func isABBA(s string) bool {
	if len(s) < 4 {
		return false
	}

	for i := 0; i < len(s)-3; i++ {
		if s[i] != s[i+1] && s[i+1] == s[i+2] && s[i] == s[i+3] {
			return true
		}
	}

	return false
}

func supportsTLS(s string) bool {

	// The idea is that, since we don't seem to have nested hypernet squences
	// we consider odd parts as normal, and even parts as hypernets
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '[' || r == ']'
	})

	valid := false
	for i, address := range parts {
		if isABBA(address) {
			// ABBA sequence in hypernet
			if i%2 != 0 {
				return false
			} else {
				valid = true
			}
		}
	}

	return valid
}

func solve1(s []string) int {
	res := 0

	for _, line := range s {
		if supportsTLS(line) {
			res++
		}
	}

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
