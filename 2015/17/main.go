package main

import (
	"flag"
	"fmt"
	"io"
	"log"
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

const target int = 150

func solve1(s []string) int {
	res := 0

	containers := []int{}
	for _, line := range s {
		containers = append(containers, ToInt(line))
	}
	dbg("%v", containers)
	for k := 1; k <= len(containers); k++ {
		for combination := range CombinationsInPlace(containers, k) {
			dbg("Combination: %v", combination)
			capacity := 0
			for _, container := range combination {
				capacity += container
			}

			if capacity == target {
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
