package main

import (
	"flag"
	"fmt"
	. "github.com/agimenez/adventofcode/utils"
	"io"
	"log"
	"maps"
	"os"
	"strings"
	"time"
	"unicode"
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

var rules map[string]int = map[string]int{
	"children": 3,
	"cats":     7,
	"samoyeds": 2,
	"akitas":   0,
	"vizslas":  0,
	"goldfish": 0,
	"trees":    3,
	"cars":     2,
	"perfumes": 1,
}

func parseAunt(l string) map[string]int {
	a := map[string]int{}

	maps.Insert(a, func(yield func(k string, v int) bool) {
		for k := range rules {
			if !yield(k, -1) {
				return
			}
		}
	})

	parts := strings.FieldsFunc(l, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	dbg("Parts: %v", parts)
	// Parts: [Sue 500 cats 2 goldfish 9 children 8]
	for i := 2; i < len(parts)-1; i += 2 {
		a[parts[i]] = ToInt(parts[i+1])
	}

	return a
}

func checkAunt(aunt map[string]int) bool {
	for k, v := range aunt {
		if v == -1 {
			continue
		}

		if rules[k] != v {
			return false
		}
	}

	return true
}

func solve1(s []string) int {
	res := 0

	for i, line := range s {
		aunt := parseAunt(line)
		dbg("Aunt %3d: %v", i+1, aunt)
		if checkAunt(aunt) {
			return i + 1
		}
	}

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
