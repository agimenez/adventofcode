package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
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

func solve1(s []string) int {
	res := 0

	replacements := map[string][]string{}
	var molecule string
	for _, line := range s {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " => ")
		if len(parts) == 2 {
			replacements[parts[0]] = append(replacements[parts[0]], parts[1])
			continue
		}

		molecule = parts[0]
	}
	dbg("%v\n%v", replacements, molecule)

	distinct := map[string]bool{}
	for from, to := range replacements {
		re := regexp.MustCompile(from)
		matches := re.FindAllStringIndex(molecule, -1)
		for _, match := range matches {
			dbg("Matched %v (=> %v) at %v", from, to, match)

			for _, repl := range to {
				var b bytes.Buffer

				// Copy until the start of the match
				b.WriteString(molecule[:match[0]])

				// Write replacement
				b.WriteString(repl)

				// Write rest of the original molecule
				b.WriteString(molecule[match[1]:])

				distinct[b.String()] = true
			}

		}

	}
	dbg("Distinct: %v", distinct)

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
