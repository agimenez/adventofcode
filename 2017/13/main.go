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

type Firewall map[int]int

// I'm sure this can be done mathemathically, but who cares when you can simulate
func solve1(s []string) int {
	res := 0

	f := Firewall{}
	maxLayer := 0
	for _, layer := range s {
		parts := strings.Split(layer, ": ")
		maxLayer = ToInt(parts[0])
		f[maxLayer] = ToInt(parts[1])
	}

	dbg("%v (%v)", f, maxLayer)
	for layer := 0; layer <= maxLayer; layer++ {
		if depth, ok := f[layer]; ok {
			cycle := 2 * (depth - 1)

			if layer%cycle == 0 {
				res += layer * depth
			}
		}
	}

	dbg("")
	return res
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")
	f := Firewall{}
	maxLayer := 0
	for _, layer := range s {
		parts := strings.Split(layer, ": ")
		maxLayer = ToInt(parts[0])
		f[maxLayer] = ToInt(parts[1])
	}

	dbg("%v (%v)", f, maxLayer)
	for ; ; res++ {
		caught := false
		for layer := 0; layer <= maxLayer; layer++ {
			if depth, ok := f[layer]; ok {
				cycle := 2 * (depth - 1)

				if (layer+res)%cycle == 0 {
					caught = true
					break

				}
			}
		}
		if !caught {
			break
		}
	}

	return res
}
