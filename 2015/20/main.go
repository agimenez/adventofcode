package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
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

// This is basically stollen from Reddit
func deliver(target int) int {
	now := time.Now()
	// houses := make(map[int]int, target/10)
	houses := make([]int, target/10)
	fmt.Printf("Slice creation time: %v\n", time.Since(now))

	for elf := 1; elf < target/10; elf++ {
		for house := elf; house < target/10; house += elf {
			houses[house] += elf * 10
			dbg("House %v got %v presents from elf %v (total: %v)", house, elf*10, elf, houses[house])
		}
		dbg("")
	}

	now = time.Now()
	lowest := slices.IndexFunc(houses, func(presents int) bool {
		return presents > target
	})
	fmt.Printf("Search time: %v\n", time.Since(now))

	return lowest
}

// Lots of duplicate but who cares
func deliver2(target int) int {
	now := time.Now()
	houses := make(map[int]int, target)
	fmt.Printf("Slice creation time: %v\n", time.Since(now))

	for elf := 1; elf < target/10; elf++ {
		for house, hcount := elf, 0; hcount < 50; house, hcount = house+elf, hcount+1 {
			houses[house] += elf * 11
			// dbg("House %v got %v presents from elf %v (total: %v)", house, elf*10, elf, houses[house])
		}

		if houses[elf] >= target {
			return elf
		}
	}

	return -1
}

func solve1(s []string) int {
	res := 0
	target := ToInt(s[0])
	res = deliver(target)

	return res
}

func solve2(s []string) int {
	res := 0
	target := ToInt(s[0])
	res = deliver2(target)

	return res
}
