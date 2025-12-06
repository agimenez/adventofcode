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

func findMax(s string, start int, end int) int {
	res := 0

	max := 0
	dbg("  -> Finding max for %v", s)
	for ; start < end; start++ {
		v := s[start]
		num := ToInt(string(v))
		dbg("    -> max: %v, idx: %v", max, start)
		dbg("    -> Checking num %c (%v)", v, num)

		if num > max {
			max = num
			res = start
			dbg("    -> NEW max: %v, idx: %v", max, res)
		}
	}

	dbg("    -> FINAL max: %v, idx: %v", max, res)
	return res
}

func solve1(s []string) int {
	res := 0

	for _, bank := range s {
		dbg("Bank: %v", bank)
		decidx := findMax(bank, 0, len(bank)-1)
		unitidx := findMax(bank, decidx+1, len(bank))

		jolts := ToInt(fmt.Sprintf("%c%c", bank[decidx], bank[unitidx]))
		dbg(" -> jolts: %v", jolts)
		res += jolts
	}

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
