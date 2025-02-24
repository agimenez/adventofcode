package main

import (
	"flag"
	"io"
	"log"
	"math"
	"os"
	"strings"
	"time"

	. "github.com/agimenez/adventofcode/utils"
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
	for _, l := range s {
		min := math.MaxInt
		max := 0

		for _, num := range strings.Fields(l) {
			n := ToInt(num)

			max = Max(max, n)
			min = Min(min, n)
		}

		res += max - min
	}

	return res
}

func divisible(v []int) int {
	res := 0

	dbg("Vals: %v", v)
	for i := range v {
		dbg(" -> %v (%v)", i, v[i])
		for j := i + 1; j < len(v); j++ {
			dbg("    -> %v (%v)", j, v[j])
			v1 := v[i]
			v2 := v[j]
			if v1%v2 == 0 {
				dbg(" RET %v", v1/v2)
				return v1 / v2
			}

			if v2%v1 == 0 {
				dbg(" RET %v", v2/v1)
				return v2 / v1
			}
		}
	}

	return res
}
func solve2(s []string) int {
	res := 0

	for _, l := range s {
		vals := []int{}
		for _, num := range strings.Fields(l) {
			vals = append(vals, ToInt(num))
		}

		res += divisible(vals)

	}
	return res
}
