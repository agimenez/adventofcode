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

	freqs := make([]map[rune]int, len(s[0]))
	dbg("Freqs: %+v", freqs)
	for _, l := range s {
		for i, c := range l {
			if freqs[i] == nil {
				freqs[i] = map[rune]int{}
			}
			freqs[i][c]++
		}
	}

	out := []rune{}
	for _, f := range freqs {
		var maxCount int
		var maxCh rune
		for ch, count := range f {
			if count > maxCount {
				maxCount = count
				maxCh = ch
			}
		}
		out = append(out, maxCh)
	}
	fmt.Printf("%s\n", string(out))

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
