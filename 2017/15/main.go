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

const div uint64 = 2147483647

type Generator struct {
	value  uint64
	factor uint64
}

func (g Generator) Value() uint64 {
	return g.value
}

func (g *Generator) Next(multiple uint64) Generator {
	next := (g.value * g.factor) % div
	for next%multiple != 0 {
		next = (next * g.factor) % div
	}

	g.value = next

	return *g
}

func solve1(s []string) int {
	res := 0

	parts := strings.Fields(s[0])
	genA := Generator{
		value:  uint64(ToInt(parts[len(parts)-1])),
		factor: 16807,
	}

	parts = strings.Fields(s[1])
	genB := Generator{
		value:  uint64(ToInt(parts[len(parts)-1])),
		factor: 48271,
	}
	for range 40_000_000 {
		vA := genA.Next(1).Value()
		vB := genB.Next(1).Value()

		if vA&0xffff == vB&0xffff {
			res++
		}
	}

	dbg("")
	return res
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")
	parts := strings.Fields(s[0])
	genA := Generator{
		value:  uint64(ToInt(parts[len(parts)-1])),
		factor: 16807,
	}

	parts = strings.Fields(s[1])
	genB := Generator{
		value:  uint64(ToInt(parts[len(parts)-1])),
		factor: 48271,
	}
	for range 5_000_000 {
		vA := genA.Next(4).Value()
		vB := genB.Next(8).Value()
		// dbg("%032b (%d)", vA, vA)
		// dbg("%032b (%d)", vB, vB)
		// dbg("")

		if vA&0xffff == vB&0xffff {
			res++
		}
	}

	return res
}
