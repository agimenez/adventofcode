package main

import (
	"flag"
	"fmt"
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

func valid(pass string) bool {
	counts := map[string]int{}

	for _, word := range strings.Fields(pass) {
		counts[word]++
		if counts[word] > 1 {
			return false
		}
	}

	return true
}

func solve1(s []string) int {
	res := 0

	for _, pass := range s {
		if valid(pass) {
			res++
		}
	}

	return res
}

func noAnagrams(pass string) bool {
	anagrams := map[string]bool{}
	for _, word := range strings.Fields(pass) {

		// Some form of anagram of this word has been seen
		if _, ok := anagrams[word]; ok {
			return false
		}

		// Record all anagrams of this word
		for comb := range Permutations([]byte(word)) {
			anagrams[string(comb)] = true
		}
	}

	return true
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")
	for _, pass := range s {
		if valid(pass) && noAnagrams(pass) {
			res++
		}
	}

	return res
}
